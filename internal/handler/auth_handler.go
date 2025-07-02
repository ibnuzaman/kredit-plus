package handler

import (
	"kredit-plus/internal/service"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service service.AuthService
}

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		service: service,
	}
}

func (a *authHandler) Login(ctx *fiber.Ctx) error {
	return ctx.SendString("Login successful")
}

func (a *authHandler) Me(ctx *fiber.Ctx) error {
	return ctx.SendString("User information retrieved successfully")
}
