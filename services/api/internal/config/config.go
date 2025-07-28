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
	GoEnv              GoEnv
	ServerUrl          string
	Port               string
	PostgresURI        string
	MinioEndpointUrl    string
	MinioAccessId      string
	MinioSecretKey     string
	MinioStorageBucket string
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
		GoEnv:              GoEnv(getEnv("GO_ENV", "development")),
		ServerUrl:          getEnv("SERVER_URL", "http://localhost:1323"),
		Port:               getEnv("PORT", "1323"),
		PostgresURI:        getEnv("POSTGRES_URI", "postgresql://root:123456@localhost:5432/postgres"),
		MinioEndpointUrl:    getEnv("MINIO_ENDPOINT_URL", "localhost:9000"),
		MinioAccessId:      getEnv("MINIO_ACCESS_ID", "minioadmin"),
		MinioSecretKey:     getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinioStorageBucket: getEnv("MINIO_STORAGE_BUCKET", "smartik"),
	}

	return config, err
}
