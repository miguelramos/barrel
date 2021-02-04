package config

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// EnvironmentConfig flags
type EnvironmentConfig struct {
	BarrelPort              string `env:"BARREL_PORT" mapstructure:"port"`
	BarrelHost              string `env:"BARREL_HOST" mapstructure:"host"`
	BarrelDatabaseURL       string `env:"BARREL_DATABASE_URL" mapstructure:"database_url"`
	BarrelDatabaseNamespace string `env:"BARREL_DATABASE_NAMESPACE" mapstructure:"database_namespace"`
	BarrelJWTSecret         string `env:"BARREL_JWT_SECRET" mapstructure:"jwt_secret"`
	BarrelBaseURL           string `env:"BARREL_BASE_URL" mapstructure:"base_url"`
	BarrelGotrueURL         string `env:"BARREL_GOTRUE_URL" mapstructure:"gotrue_url"`
	BarrelMinioAccessKey    string `env:"BARREL_MINIO_ACCESS_KEY" mapstructure:"minio_access_key"`
	BarrelMinioSecretKey    string `env:"BARREL_MINIO_SECRET_KEY" mapstructure:"minio_secret_key"`
	BarrelMinioPort         string `env:"BARREL_MINIO_PORT" mapstructure:"minio_port"`
	BarrelMinioURL          string `env:"BARREL_MINIO_URL" mapstructure:"minio_url"`
	BarrelRolesPath         string `env:"BARREL_ADMIN_ROLES_PATH" mapstructure:"admin_roles_path"`
	BarrelAdminRole         string `env:"BARREL_ADMIN_ROLE" mapstructure:"admin_role"`
	BarrelAdminKey          string `env:"BARREL_ADMIN_KEY" mapstructure:"admin_key"`
}

// LoadEnvironmentConfig load config from env
func LoadEnvironmentConfig() *EnvironmentConfig {
	viper := viper.New()

	loadEnv(viper)

	envConfig := new(EnvironmentConfig)

	viper.Unmarshal(&envConfig)

	return envConfig
}

func loadDefault(viper *viper.Viper) {
	viper.SetDefault("BARREL_MINIO_PORT", "9876")
	viper.SetDefault("BARREL_MINIO_URL", "localhost:9876")
	viper.SetDefault("BARREL_PORT", "9078")
	viper.SetDefault("BARREL_HOST", "localhost")
	viper.SetDefault("BARREL_DATABASE_NAMESPACE", "storage")
	viper.SetDefault("BARREL_BASE_URL", "http://localhost:9078")
	viper.SetDefault("BARREL_GOTRUE_URL", "http://localhost:9999")
}

func loadEnv(viper *viper.Viper) {
	dir, _ := os.Getwd()
	envfile := path.Join(dir, ".env")

	viper.SetEnvPrefix("barrel")
	viper.SetConfigFile(envfile)

	viper.AutomaticEnv()
	viper.BindEnv("MINIO_PORT")
	viper.BindEnv("MINIO_URL")
	viper.BindEnv("MINIO_ACCESS_KEY")
	viper.BindEnv("MINIO_SECRET_KEY")
	viper.BindEnv("PORT")
	viper.BindEnv("HOST")
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("DATABASE_NAMESPACE")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("BASE_URL")
	viper.BindEnv("GOTRUE_URL")
	viper.BindEnv("ADMIN_ROLES_PATH")
	viper.BindEnv("ADMIN_ROLE")
	viper.BindEnv("ADMIN_KEY")

	loadDefault(viper)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Info("Env file not present")
	}
}
