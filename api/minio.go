package api

import "github.com/gofiber/fiber/v2"

func (api *API) CreateUser(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "identity",
	})
}
