package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/utils"
)

// CreateUser create user on minio
func (api *API) CreateUser(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)

	if isPrivate && !isAdmin {
		return utils.NewException(utils.ErrorUserCreation, fiber.StatusForbidden, "Creation permission denied")
	}

	identity := new(config.Identity)

	if err := ctx.BodyParser(identity); err != nil {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, "Invalid request body parser")
	}

	if err := identity.Validate(); len(err.Errors) > 0 {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, err.Error())
	}

	if err := config.CreateOrgUser(api.config, identity.AccessKey, identity.SecretKey); err != nil {
		return utils.NewException(utils.ErrorOrgUserFailure, fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": identity,
	})
}
