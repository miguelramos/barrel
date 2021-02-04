package config

import (
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	madmin "github.com/minio/minio/pkg/madmin"
	"github.com/stretchr/testify/assert"
)

func TestMinioOpenClient(t *testing.T) {
	os.Setenv("BARREL_MINIO_PORT", "9876")
	os.Setenv("BARREL_MINIO_URL", "localhost:9876")
	os.Setenv("BARREL_MINIO_ACCESS_KEY", "A5a0a87b725552daXd")
	os.Setenv("BARREL_MINIO_SECRET_KEY", "369f5e7b4a41e25452c353D629a24c372b62c90")
	os.Setenv("BARREL_PORT", "9078")
	os.Setenv("BARREL_HOST", "localhost")
	os.Setenv("BARREL_DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	os.Setenv("BARREL_DATABASE_NAMESPACE", "storage")
	os.Setenv("BARREL_BASE_URL", "http://localhost:9078")

	cfg := LoadEnvironmentConfig()

	client, err := OpenClient(cfg)
	admin, err := OpenAdminClient(cfg)
	clientNew, err := NewClient(cfg, "", "", "")

	assert.IsType(t, new(minio.Client), client)
	assert.IsType(t, new(minio.Client), clientNew)
	assert.IsType(t, new(madmin.AdminClient), admin)

	assert.Nil(t, err)
}
