package middleware

import (
	"kredit-plus/config"
	"kredit-plus/internal/repository"
	"kredit-plus/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Middleware interface {
	Auth() fiber.Handler
}

type middleware struct {
	authRepo repository.AuthRepository
	conf     *config.Config
	log      *zerolog.Logger
}

func NewMiddleware(authRepo repository.AuthRepository) Middleware {
	return &middleware{
		authRepo: authRepo,
		conf:     config.Get(),
		log:      logger.Get("middleware"),
	}
}
