package internal

import (
	"context"
	"database/sql"
	"fmt"
	"kredit-plus/config"
	"kredit-plus/database"
	_ "kredit-plus/docs"
	"kredit-plus/internal/handler"
	"kredit-plus/internal/middleware"
	"kredit-plus/internal/repository"
	"kredit-plus/internal/service"
	"kredit-plus/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
)

func Run() {
	logs := logger.Get("app")
	logs.Info().Msg("Application is running!")

	db := database.Get()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to close database connection")
		}
	}(db)

	dbGorm := database.GetGorm()
	conf := config.Get()

	logs.Info().Msg("Configuring server...")
	logApi := logger.Get("api")
	fiber.SetParserDecoder(fiberParserConfig())
	app := fiber.New(fiberConfig(conf))
	app.Use(fiberzerolog.New(zerologConfig(logApi)))
	app.Use(requestid.New())
	app.Use(cors.New(corsConfig(conf)))
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	app.Get("/swagger/*", basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.Swagger.Username: conf.Swagger.Password,
		},
	}), swagger.New(swagger.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  true,
		DocExpansion: "list",
	}))

	logs.Info().Msg("Server is running!")
	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	server := &http.Server{
		Addr: baseUrl,
	}

	go func() {
		err := app.Listen(baseUrl)
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Repository
	authRepo := repository.NewAuthRepository(dbGorm)

	// Service
	authService := service.NewAuthService(authRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	homeHandler := handler.NewHomeHandler()

	// Router
	mid := middleware.NewMiddleware(authRepo)
	route := NewRouter(app, mid)
	route.Home(homeHandler)
	route.Auth(authHandler)

	handleShutdown(server, logs)
}

func handleShutdown(server *http.Server, logs *zerolog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Warn().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.Shutdown(ctx); err != nil {
		logs.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logs.Info().Msg("Server exiting")
}
