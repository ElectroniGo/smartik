package config

import (
	"os"

	"github.com/joho/godotenv"
)

type GoEnv string

const (
	GoEnvDevelopment GoEnv = "development"
	GoEnvProduction  GoEnv = "production"
)

type Env struct {
	GoEnv       GoEnv
	Port        string
	PostgresURI string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func Load() (*Env, error) {
	err := godotenv.Load()

	config := &Env{
		GoEnv:       GoEnv(getEnv("GO_ENV", "development")),
		Port:        getEnv("PORT", "1323"),
		PostgresURI: getEnv("POSTGRES_URI", "postgresql://root:123456@localhost:5432/postgres"),
	}

	return config, err
}
