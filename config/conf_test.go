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
	os.Setenv("MINIO_PORT", "9876")
	os.Setenv("MINIO_URL", "localhost:9876")
	os.Setenv("MINIO_ACCESS_KEY", "A5a0a87b725552daXd")
	os.Setenv("MINIO_SECRET_KEY", "369f5e7b4a41e25452c353D629a24c372b62c90")
	os.Setenv("MINION_PORT", "9078")
	os.Setenv("MINION_HOST", "localhost")
	os.Setenv("MINION_DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	os.Setenv("MINION_DATABASE_NAMESPACE", "storage")
	os.Setenv("MINION_BASE_URL", "http://localhost:9078")

	cfg := LoadEnvironmentConfig()

	assert.Equal(t, os.Getenv("MINIO_PORT"), cfg.MinioPort)
	assert.Equal(t, os.Getenv("MINIO_URL"), cfg.MinioURL)
	assert.Equal(t, os.Getenv("MINIO_ACCESS_KEY"), cfg.MinioAccessKey)
	assert.Equal(t, os.Getenv("MINIO_SECRET_KEY"), cfg.MinioSecretKey)
	assert.Equal(t, os.Getenv("MINION_PORT"), cfg.BarrelPort)
	assert.Equal(t, os.Getenv("MINION_HOST"), cfg.BarrelHost)
	assert.Equal(t, os.Getenv("MINION_DATABASE_URL"), cfg.BarrelDatabaseURL)
	assert.Equal(t, os.Getenv("MINION_BASE_URL"), cfg.BarrelBaseURL)
}
