package api

import "github.com/gofiber/fiber/v2"

func (api *API) HealthCheck(ctx *fiber.Ctx) error {
	ctx.Status(200)

	return ctx.JSON(fiber.Map{
		"data": "Api running",
	})
}
