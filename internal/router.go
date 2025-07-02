package internal

import (
	"kredit-plus/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	app *fiber.App
}

type Router interface {
	Auth(handler handler.AuthHandler)
}

func NewRouter(app *fiber.App) Router {
	return &router{
		app: app,
	}
}

func (r *router) Auth(handler handler.AuthHandler) {
	auth := r.app.Group("v1/auth")
	auth.Post("login", handler.Login)

	// TODO: add middleware for authentication
	auth.Get("me", handler.Me)
}
