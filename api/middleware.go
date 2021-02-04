package api

import (
	"encoding/json"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/utils"
)

// AuthorizedMiddleware checks jwt authorization
func (api *API) AuthorizedMiddleware(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	authLength := len(auth)
	authBearer := len("Bearer")

	if authLength == 0 {
		return utils.NewException(utils.ErrorOrgStatusForbidden, fiber.StatusForbidden, "Only authorized requests are permitted")
	}

	if authLength > authBearer+1 && auth[:authBearer] == "Bearer" {
		bearer := auth[authBearer+1:]

		parser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}

		token, err := parser.ParseWithClaims(bearer, &config.GoTrueClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(api.config.BarrelJWTSecret), nil
		})

		if err != nil {
			return utils.NewException(utils.ErrorOrgStatusForbidden, fiber.StatusForbidden, err.Error())
		}

		claims, ok := token.Claims.(*config.GoTrueClaims)

		if !ok {
			return utils.NewException(utils.ErrorOrgInvalidToken, fiber.StatusNotAcceptable, "Your token is not valid")
		}

		ctx.Locals("claims", claims)
		ctx.Locals("token", auth)
	} else {
		return utils.NewException(utils.ErrorOrgStatusForbidden, fiber.StatusForbidden, "Only authorized requests are permitted")
	}

	return ctx.Next()
}

// AdminMiddleware verify and set if user is admin
func (api *API) AdminMiddleware(ctx *fiber.Ctx) error {
	var roles []string
	var isAdmin bool = false

	headerKey := ctx.Get("X-BARREL-KEY")

	if len(headerKey) > 0 && strings.Compare(headerKey, api.config.BarrelAdminKey) == 0 {
		isAdmin = true
	} else {
		claimer := ctx.Locals("claims").(*config.GoTrueClaims)
		claimJSON, err := json.Marshal(claimer)

		if err == nil && gjson.ValidBytes(claimJSON) {
			result := gjson.GetBytes(claimJSON, api.config.BarrelRolesPath)

			if result.Index > 0 {
				err := json.Unmarshal([]byte(result.Raw), &roles)

				if err == nil {
					isAdmin = utils.Contains(roles, api.config.BarrelAdminRole)
				}
			}
		}
	}

	ctx.Locals("admin", isAdmin)
	logrus.Infof("Is admin: %v", isAdmin)

	return ctx.Next()
}

func (api *API) CanAccessMiddleware(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("admin").(bool)

	if !isAdmin {
		config.UserIsRegister(api.config)
	}

	return ctx.Next()
}
