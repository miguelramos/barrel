package config

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	iampolicy "github.com/minio/minio/pkg/iam/policy"
	madmin "github.com/minio/minio/pkg/madmin"
	"github.com/websublime/barrel/utils"
)

// OpenClient open minio client connection
func OpenClient(conf *EnvironmentConfig) (*minio.Client, error) {
	minioClient, err := minio.New(conf.BarrelMinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.BarrelMinioAccessKey, conf.BarrelMinioSecretKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

// OpenAdminClient open minio admin client
func OpenAdminClient(conf *EnvironmentConfig) (*madmin.AdminClient, error) {
	minioAdmin, err := madmin.New(conf.BarrelMinioURL, conf.BarrelMinioAccessKey, conf.BarrelMinioSecretKey, false)

	if err != nil {
		return nil, err
	}

	return minioAdmin, nil
}

// NewClient create a new client connection
func NewClient(conf *EnvironmentConfig, key string, secret string, token string) (*minio.Client, error) {
	minioClient, err := minio.New(conf.BarrelMinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(key, secret, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func UserIsRegister(conf *EnvironmentConfig, key string) (*madmin.UserInfo, error) {
	admin, _ := OpenAdminClient(conf)

	users, err := admin.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}

	user, ok := users[key]
	if !ok {
		return nil, utils.NewException(utils.ErrorOrgStatusForbidden, fiber.StatusForbidden, "User not found")
	}

	return &user, nil
}

func CreateOrgUser(conf *EnvironmentConfig, accessKey string, secretKey string) error {
	// https://github.com/minio/minio/blob/master/pkg/madmin/examples/add-user-and-policy.go
	admin, _ := OpenAdminClient(conf)

	if err := admin.AddUser(context.Background(), accessKey, secretKey); err != nil {
		return err
	}

	return nil
}

func CreateCannedPolicy(conf *EnvironmentConfig, name string, policy *iampolicy.Policy) error {
	admin, _ := OpenAdminClient(conf)

	if err := admin.AddCannedPolicy(context.Background(), name, policy); err != nil {
		return err
	}

	return nil
}

func CreateUserPolicy(conf *EnvironmentConfig, key string, policy string) error {
	admin, _ := OpenAdminClient(conf)

	if err := admin.SetPolicy(context.Background(), policy, key, false); err != nil {
		return err
	}

	return nil
}
