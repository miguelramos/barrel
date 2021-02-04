package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/models"
	"github.com/websublime/barrel/utils"
)

func (api *API) CreateBucket(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)

	if isPrivate && !isAdmin {
		return utils.NewException(utils.ErrorBucketCreationForbidden, fiber.StatusForbidden, "Creation permission denied")
	}

	bucket, err := models.NewBucket("", isPrivate)
	if err != nil {
		return utils.NewException(utils.ErrorBucketModel, fiber.StatusUnprocessableEntity, "Unable to create bucket model")
	}

	if err := ctx.BodyParser(bucket); err != nil {
		return utils.NewException(utils.ErrorBucketBodyParse, fiber.StatusPreconditionFailed, "Invalid request body parser")
	}

	return ctx.JSON(fiber.Map{
		"data": "identity",
	})
}
