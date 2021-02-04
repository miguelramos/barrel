package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/storage"
)

type API struct {
	db     *storage.Connection
	config *config.EnvironmentConfig
	store  *minio.Client
	app    *fiber.App
}

func WithVersion(app *fiber.App, conf *config.EnvironmentConfig, db *storage.Connection, store *minio.Client) {
	api := &API{
		db:     db,
		config: conf,
		store:  store,
		app:    app,
	}

	publicRouter := app.Group("/v1")
	privateRouter := publicRouter.Group("/org")

	NewPublicApi(api, publicRouter)
	NewPrivateApi(api, privateRouter)
}

func NewPublicApi(api *API, router fiber.Router) {
	router.Get("", api.HealthCheck)
}

func NewPrivateApi(api *API, router fiber.Router) {
	// TODO: private
}
