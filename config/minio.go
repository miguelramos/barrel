package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	madmin "github.com/minio/minio/pkg/madmin"
)

func OpenClient(conf *EnvironmentConfig) *minio.Client {
	minioClient, err := minio.New(conf.MinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinioAccessKey, conf.MinioSecretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

func OpenAdminClient(conf *EnvironmentConfig) (*madmin.AdminClient, error) {
	minioAdmin, err := madmin.New(conf.MinioURL, conf.MinioAccessKey, conf.MinioSecretKey, false)

	if err != nil {
		return nil, err
	}

	return minioAdmin, nil
}

func NewClient(conf *EnvironmentConfig, key string, secret string, token string) (*minio.Client, error) {
	minioClient, err := minio.New(conf.MinioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(key, secret, token),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func CreateOrgPolicy(conf *EnvironmentConfig) {
	// https://github.com/minio/minio/blob/master/pkg/madmin/examples/add-user-and-policy.go
}
