package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/websublime/barrel/api"
	"github.com/websublime/barrel/config"
	"github.com/websublime/barrel/storage"
)

func main() {
	boot()
}

func boot() {
	env := config.LoadEnvironmentConfig()

	db, err := storage.Dial(env)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()

	app := config.BootApplication()

	client, err := config.OpenClient(env)
	if err != nil {
		logrus.Fatalf("Error opening minio: %+v", err)
	}

	api.WithVersion(app, env, db, client)

	app.Listen(fmt.Sprintf("%s:%s", env.BarrelHost, env.BarrelPort))
}
