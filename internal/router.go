package internal

import (
	"kredit-plus/internal/handler"
	"kredit-plus/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	app *fiber.App
	mid middleware.Middleware
}

type Router interface {
	Auth(handler handler.AuthHandler)
}

func NewRouter(app *fiber.App, mid middleware.Middleware) Router {
	return &router{
		app: app,
		mid: mid,
	}
}

func (r *router) Auth(handler handler.AuthHandler) {
	auth := r.app.Group("v1/auth")
	auth.Post("login", handler.Login)

	// TODO: add middleware for authentication
	auth.Get("me", r.mid.Auth(), handler.Me)
}
