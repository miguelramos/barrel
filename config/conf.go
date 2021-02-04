package config

import "github.com/spf13/viper"

type EnvironmentConfig struct {
	BarrelPort              string
	BarrelHost              string
	BarrelDatabaseURL       string
	BarrelDatabaseNamespace string
	BarrelJWTSecret         string
	BarrelBaseURL           string
	BarrelGotrueURL         string
	MinioAccessKey          string
	MinioSecretKey          string
	MinioPort               string
	MinioURL                string
}

func LoadEnvironmentConfig() *EnvironmentConfig {
	viper := viper.New()

	loadEnv(viper)
	loadDefault(viper)

	envConfig := new(EnvironmentConfig)

	viper.Unmarshal(&envConfig)

	return envConfig
}

func loadDefault(viper *viper.Viper) {
	viper.SetDefault("MINIO_PORT", "9876")
	viper.SetDefault("MINIO_URL", "localhost:9876")
	viper.SetDefault("MINION_PORT", "9078")
	viper.SetDefault("MINION_HOST", "localhost")
	viper.SetDefault("MINION_DATABASE_NAMESPACE", "storage")
	viper.SetDefault("MINION_BASE_URL", "http://localhost:9078")
	viper.SetDefault("MINION_GOTRUE_URL", "http://localhost:9999")
}

func loadEnv(viper *viper.Viper) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.WatchConfig()
}
