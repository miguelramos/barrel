package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
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

// PutObject upload to bucket
func (api *API) PutObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket")
	medias := []*models.Media{}
	//TODO: user
	bucket, err := models.FindBucket(api.db, bucketName)
	if err != nil {
		return utils.NewException(utils.ErrorBucketMissing, fiber.StatusBadRequest, err.Error())
	}

	if et, err := exiftool.NewExiftool(); err == nil {
		defer et.Close()

		if form, err := ctx.MultipartForm(); err == nil {
			files := form.File["asset"]

			for _, file := range files {
				ctx.SaveFile(file, fmt.Sprintf("./temp/%s", file.Filename))
				data := et.ExtractMetadata(fmt.Sprintf("./temp/%s", file.Filename))

				meta, err := json.Marshal(data[0].Fields)
				if err != nil {
					return utils.NewException(utils.ErrorResourceMetaFailure, fiber.StatusBadRequest, err.Error())
				}

				metafile := new(config.MetaFile)
				json.Unmarshal([]byte(meta), &metafile)

				bucketFile, err := api.store.FPutObject(ctx.Context(), bucketName, metafile.FileName, fmt.Sprintf("./temp/%s", metafile.FileName), minio.PutObjectOptions{
					ContentType:  metafile.MIMEType,
					UserMetadata: map[string]string{},
				})

				if err != nil {
					return utils.NewException(utils.ErrorResourceBucketFailure, fiber.StatusBadRequest, err.Error())
				}

				u, _ := url.Parse(api.config.BarrelBaseURL)
				u.Path = path.Join(u.Path, bucketName, metafile.FileName)
				bucketFileJSON, _ := json.Marshal(bucketFile)
				metaFileJSON, _ := json.Marshal(metafile)

				media, _ := models.NewMedia(u.String(), nulls.NewUUID(uuid.Nil), nulls.NewString(string(bucketFileJSON)), nulls.NewString(string(metaFileJSON)))
				bucketMedia, _ := models.NewBucketMedia(bucket.ID, media.ID)

				err = api.db.Transaction(func(tx *storage.Connection) error {
					terr := tx.Create(media)
					terr = tx.Create(bucketMedia)

					return terr
				})
				if err != nil {
					return utils.NewException(utils.ErrorResourceModelSave, fiber.StatusBadRequest, err.Error())
				}

				os.Remove(fmt.Sprintf("./temp/%s", file.Filename))

				medias = append(medias, media)
			}

			return ctx.JSON(fiber.Map{
				"data": medias,
			})
		} else {
			return utils.NewException(utils.ErrorResourceInvalidForm, fiber.StatusBadRequest, err.Error())
		}
	} else {
		return utils.NewException(utils.ErrorExifMissing, fiber.StatusBadRequest, err.Error())
	}

}
