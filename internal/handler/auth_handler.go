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

// Login godoc
//
//	@Summary		Authenticate user login
//	@Description	Handles user login by verifying credentials and returning authentication data.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.LoginRequest	true	"Payload"
//	@Success		200		{object}	model.BaseResponse	"Successful login response"
//	@Failure		401		{object}	model.BaseResponse	"Unauthorized error response"
//	@Router			/v1/auth/login [post]
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

// Me godoc
//
//	@Summary		Get authenticated user info
//	@Description	Retrieves information about the currently authenticated user.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Success		200	{object}	model.BaseResponse	"Successful user info response"
//	@Failure		401	{object}	model.BaseResponse	"Unauthorized error response"
//	@Router			/v1/auth/me [get]
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
