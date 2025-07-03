package handler

import (
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/service"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	service   service.CustomerService
	exception exception.Exception
}

type CustomerHandler interface {
	Information(ctx *fiber.Ctx) error
	Tenor(ctx *fiber.Ctx) error
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return &customerHandler{
		service:   service,
		exception: exception.NewException(),
	}
}

// Information godoc
//
//	@Summary		Get customer information
//	@Description	Retrieves information about the customer.
//	@Tags			customer
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Success		200	{object}	model.BaseResponse	"Successful customer information response"
//	@Failure		401	{object}	model.BaseResponse	"Unauthorized error response"
//	@Router			/v1/customer/information [get]
func (h *customerHandler) Information(ctx *fiber.Ctx) error {
	user := new(model.AuthMe)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	data := h.service.Information(ctx.UserContext(), user.ID)
	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieved customer information",
		Data:    data,
	})
}

func (h *customerHandler) Tenor(ctx *fiber.Ctx) error {
	return nil
}
