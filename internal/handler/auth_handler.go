package handler

import (
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/service"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service   service.AuthService
	exception exception.Exception
}

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		service:   service,
		exception: exception.NewException(),
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	var req model.LoginRequest
	err := ctx.BodyParser(&req)
	h.exception.BadRequestErr(err)
	data := h.service.Login(ctx.Context(), req)
	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success verify login by google",
		Data:    data,
	})
}

func (h *authHandler) Me(ctx *fiber.Ctx) error {
	user := new(model.AuthMe)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success get user info",
		Data:    user,
	})
}
