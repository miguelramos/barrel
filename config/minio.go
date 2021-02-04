package config

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	madmin "github.com/minio/minio/pkg/madmin"
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
		Creds:  credentials.NewStaticV4(key, secret, token),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func UserIsRegister(conf *EnvironmentConfig) {
	admin, _ := OpenAdminClient(conf)

	users, _ := admin.ListUsers(context.Background())

	fmt.Print(users)
}

func CreateOrgPolicy(conf *EnvironmentConfig) {
	// https://github.com/minio/minio/blob/master/pkg/madmin/examples/add-user-and-policy.go
}
