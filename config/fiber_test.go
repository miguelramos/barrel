package config

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberInstance(t *testing.T) {
	app := BootApplication()

	assert.IsType(t, new(fiber.App), app)
}
