package model

import (
	"kredit-plus/constant"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AuthMe struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *AuthMe) FromReq(ctx *fiber.Ctx) bool {
	identity, isFound := ctx.Locals(constant.KeyLocalsAuthUser).(AuthMe)
	if isFound {
		*a = identity
	}
	return isFound
}
