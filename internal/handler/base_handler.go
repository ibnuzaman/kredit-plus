package handler

import (
	"kredit-plus/constant"

	"github.com/gofiber/fiber/v2"
)

func getPaginator(ctx *fiber.Ctx) (uint, uint) {
	page := ctx.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	perPage := ctx.QueryInt("per_page", 10)
	if perPage < 1 {
		perPage = constant.DefaultPerPage
	}

	return uint(page), uint(perPage)
}
