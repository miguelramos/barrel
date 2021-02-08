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

// CreateCannedPolicy create a canned policy
func (api *API) CreateCannedPolicy(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)

	if isPrivate && !isAdmin {
		return utils.NewException(utils.ErrorUserCreation, fiber.StatusForbidden, "Creation permission denied")
	}

	canned := new(config.CannedPolicy)

	if err := ctx.BodyParser(canned); err != nil {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, "Invalid request body parser")
	}

	if err := canned.Validate(); len(err.Errors) > 0 {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, err.Error())
	}

	policyType := GetPolicyType(canned.Policy)
	policy, _, err := CreatePolicy(policyType, canned.Bucket)
	if err != nil {
		return utils.NewException(utils.ErrorOrgPolicyFailure, fiber.StatusPreconditionFailed, err.Error())
	}

	if err := config.CreateCannedPolicy(api.config, canned.Name, policy); err != nil {
		return utils.NewException(utils.ErrorOrgPolicyCreate, fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": canned,
	})
}

// SetPolicy set policy iampolicy
func (api *API) SetPolicy(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)

	if isPrivate && !isAdmin {
		return utils.NewException(utils.ErrorUserCreation, fiber.StatusForbidden, "Creation permission denied")
	}

	canned := new(config.CannedPolicy)

	if err := ctx.BodyParser(canned); err != nil {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, "Invalid request body parser")
	}

	if err := canned.ValidatePolicy(); len(err.Errors) > 0 {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, err.Error())
	}

	policyType := GetPolicyType(canned.Policy)
	policy, _, err := CreatePolicy(policyType, canned.Bucket)
	if err != nil {
		return utils.NewException(utils.ErrorOrgPolicyFailure, fiber.StatusPreconditionFailed, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": policy,
	})
}

//CreateUserPolicy associate a canned policy to a user
func (api *API) CreateUserPolicy(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)

	if isPrivate && !isAdmin {
		return utils.NewException(utils.ErrorUserCreation, fiber.StatusForbidden, "Creation permission denied")
	}

	policy := new(config.IdentityPolicy)

	if err := ctx.BodyParser(policy); err != nil {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, "Invalid request body parser")
	}

	if err := policy.Validate(); len(err.Errors) > 0 {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusPreconditionFailed, err.Error())
	}

	if err := config.CreateUserPolicy(api.config, policy.AccessKey, policy.PolicyName); err != nil {
		return utils.NewException(utils.ErrorUserBodyParse, fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": policy,
	})
}
