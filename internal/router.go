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
	Home(handler handler.HomeHandler)
	Auth(handler handler.AuthHandler)
	Customer(handler handler.CustomerHandler)
	Loan(handler handler.LoanHandler)
	Transaction(handler handler.TransactionHandler)
}

func NewRouter(app *fiber.App, mid middleware.Middleware) Router {
	return &router{
		app: app,
		mid: mid,
	}
}

func (r *router) Home(handler handler.HomeHandler) {
	r.app.Get("", handler.Home)
}

func (r *router) Auth(handler handler.AuthHandler) {
	auth := r.app.Group("v1/auth")
	auth.Post("login", handler.Login)
	auth.Get("me", r.mid.Auth(), handler.Me)
}

func (r *router) Customer(handler handler.CustomerHandler) {
	customer := r.app.Group("v1/customer")
	customer.Get("information", r.mid.Auth(), handler.Information)
	customer.Get("tenor", r.mid.Auth(), handler.Tenor)
}

func (r *router) Loan(handler handler.LoanHandler) {
	loan := r.app.Group("v1/loan")
	loan.Get("", r.mid.Auth(), handler.List)
	loan.Post("", r.mid.Auth(), handler.Create)
}

func (r *router) Transaction(handler handler.TransactionHandler) {
	transaction := r.app.Group("v1/transaction")
	transaction.Get("", r.mid.Auth(), handler.List)
	transaction.Post("", r.mid.Auth(), handler.Create)
}
