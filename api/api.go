package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/storage"
)

// API api instance
type API struct {
	db     *storage.Connection
	config *config.EnvironmentConfig
	store  *minio.Client
	app    *fiber.App
}

// WithVersion api with version prefix
func WithVersion(app *fiber.App, conf *config.EnvironmentConfig, db *storage.Connection, store *minio.Client) {
	api := &API{
		db:     db,
		config: conf,
		store:  store,
		app:    app,
	}

	publicRouter := app.Group("/v1")
	privateRouter := publicRouter.Group("/org", api.AuthorizedMiddleware, api.AdminMiddleware, api.CanAccessMiddleware)

	NewPublicAPI(api, publicRouter)
	NewPrivateAPI(api, privateRouter)
}

// NewPublicAPI public routes
func NewPublicAPI(api *API, router fiber.Router) {
	router.Get("", api.HealthCheck)
}

// NewPrivateAPI private routes
func NewPrivateAPI(api *API, router fiber.Router) {
	router.Get("", api.HealthCheck)

	bucketRouter := router.Group("/bucket")
	bucketRouter.Post("", api.CreateBucket)

	userRouter := router.Group("/user")
	userRouter.Post("", api.CreateUser)

	policyRouter := router.Group("/policy")
	policyRouter.Post("", api.CreateCannedPolicy)
}
