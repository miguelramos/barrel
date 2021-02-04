package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	defer os.Clearenv()
	os.Exit(m.Run())
}

func TestEnvironmentConfig(t *testing.T) {
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

	assert.Equal(t, os.Getenv("BARREL_MINIO_PORT"), cfg.BarrelMinioPort)
	assert.Equal(t, os.Getenv("BARREL_MINIO_URL"), cfg.BarrelMinioURL)
	assert.Equal(t, os.Getenv("BARREL_MINIO_ACCESS_KEY"), cfg.BarrelMinioAccessKey)
	assert.Equal(t, os.Getenv("BARREL_MINIO_SECRET_KEY"), cfg.BarrelMinioSecretKey)
	assert.Equal(t, os.Getenv("BARREL_PORT"), cfg.BarrelPort)
	assert.Equal(t, os.Getenv("BARREL_HOST"), cfg.BarrelHost)
	assert.Equal(t, os.Getenv("BARREL_DATABASE_URL"), cfg.BarrelDatabaseURL)
	assert.Equal(t, os.Getenv("BARREL_BASE_URL"), cfg.BarrelBaseURL)
}
