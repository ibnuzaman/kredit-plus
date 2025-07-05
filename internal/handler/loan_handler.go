package handler

import (
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/service"

	"github.com/gofiber/fiber/v2"
)

type loanHandler struct {
	service   service.LoanService
	exception exception.Exception
}

type LoanHandler interface {
	List(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

func NewLoanHandler(service service.LoanService) LoanHandler {
	return &loanHandler{
		service:   service,
		exception: exception.NewException(),
	}
}

// List godoc
//
//	@Summary		Get loan
//	@Description	Retrieves available loans.
//	@Tags			loan
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page		query		int												false	"Page"		default(1)
//	@Param			per_page	query		int												false	"Per Page"	default(10)
//	@Success		200			{object}	model.BaseResponse{data=[]model.LoanResponse}	"Successful loan list response"
//	@Failure		401			{object}	model.BaseResponse								"Unauthorized error response"
//	@Router			/v1/loan [get]
func (h *loanHandler) List(ctx *fiber.Ctx) error {
	user := new(model.AuthMe)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	page, perPage := getPaginator(ctx)
	data := h.service.List(ctx.UserContext(), user.ID, page, perPage)
	return ctx.JSON(model.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieved loan customer",
		Data:    data,
	})
}

// Create godoc
//
//	@Summary		Create loan
//	@Description	Creates a new loan.
//	@Tags			loan
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			request	body		model.CreateLoanRequest	true	"Create Loan Request"
//	@Success		201		{object}	model.BaseResponse		"Successful loan creation response"
//	@Failure		400		{object}	model.BaseResponse		"Bad request error response"
//	@Failure		401		{object}	model.BaseResponse		"Unauthorized error response"
//	@Failure		422		{object}	model.BaseResponse		"Validation error response"
//	@Router			/v1/loan [post]
func (h *loanHandler) Create(ctx *fiber.Ctx) error {
	user := new(model.AuthMe)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	req := new(model.CreateLoanRequest)
	err := ctx.BodyParser(&req)
	h.exception.BadRequestErr(err)
	h.exception.ValidateStruct(req)

	h.service.Create(ctx.UserContext(), user.ID, *req)
	return ctx.Status(fiber.StatusCreated).JSON(model.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "Successfully created loan",
	})
}
