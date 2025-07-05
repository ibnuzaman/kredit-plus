package handler

import (
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/service"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	service   service.TransactionService
	exception exception.Exception
}

type TransactionHandler interface {
	List(ctx *fiber.Ctx) error
}

func NewTransactionHandler(service service.TransactionService) TransactionHandler {
	return &transactionHandler{
		service:   service,
		exception: exception.NewException(),
	}
}

// List godoc
//
//	@Summary		Get transaction
//	@Description	Retrieves available transactions.
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page		query		int												false	"Page"		default(1)
//	@Param			per_page	query		int												false	"Per Page"	default(10)
//	@Success		200			{object}	model.BaseResponse{data=[]model.Transaction}	"Successful transaction list response"
//	@Failure		401			{object}	model.BaseResponse								"Unauthorized error response"
//	@Router			/v1/transaction [get]
func (h *transactionHandler) List(ctx *fiber.Ctx) error {
	user := new(model.AuthMe)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	page, perPage := getPaginator(ctx)
	data := h.service.Transaction(ctx.UserContext(), user.ID, page, perPage)
	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieved transaction customer",
		Data:    data,
	})
}
