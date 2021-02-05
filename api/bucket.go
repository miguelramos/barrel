package api

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/minio/minio-go/v7"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/models"
	"github.com/websublime/barrel/storage"
	"github.com/websublime/barrel/utils"
)

//CreateBucket creates public/private bucket
func (api *API) CreateBucket(ctx *fiber.Ctx) error {
	claimer := ctx.Locals("claims").(*config.GoTrueClaims)
	isPrivate := claimer != nil
	isAdmin := ctx.Locals("admin").(bool)
	policyType := AnonymousPolicy
	bucketName := new(strings.Builder)

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

	if err := bucket.Validate(); len(err.Errors) > 0 {
		return utils.NewException(utils.ErrorBucketModel, fiber.StatusUnprocessableEntity, err.Error())
	}

	if len(bucket.Policy) > 0 {
		policyType = GetPolicyType(bucket.Policy)
	} else {
		bucket.Policy = string(policyType)
	}

	if isPrivate {
		bucketName.WriteString(slug.Make(claimer.Audience))
		bucketName.WriteString("-")
		bucketName.WriteString(slug.Make(bucket.Bucket.String))
	} else {
		bucketName.WriteString(slug.Make(bucket.Bucket.String))
	}

	if err := api.store.MakeBucket(ctx.Context(), bucketName.String(), minio.MakeBucketOptions{Region: "EUROPE-WEST2"}); err != nil {
		exists, errBucketExists := api.store.BucketExists(ctx.Context(), bucketName.String())

		if errBucketExists == nil && exists {
			return utils.NewException(utils.ErrorBucketExist, fiber.StatusBadRequest, fmt.Sprintf("Bucket %s exists", bucket.Name))
		}

		return utils.NewException(utils.ErrorBucketCreation, fiber.StatusBadRequest, err.Error())
	}

	bucket.Bucket = nulls.NewString(bucketName.String())

	err = api.db.Transaction(func(tx *storage.Connection) error {
		terr := tx.Create(bucket)

		return terr
	})
	if err != nil {
		return utils.NewException(utils.ErrorBucketModelTransaction, fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": bucket,
	})
}

func (api *API) PutObject(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "object",
	})
}
