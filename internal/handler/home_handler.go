package handler

import (
	"kredit-plus/internal/model"

	"github.com/gofiber/fiber/v2"
)

type homeHandler struct{}

type HomeHandler interface {
	Home(ctx *fiber.Ctx) error
}

func NewHomeHandler() HomeHandler {
	return &homeHandler{}
}

func (h *homeHandler) Home(ctx *fiber.Ctx) error {
	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Welcome to Kredit Plus API",
	})
}
