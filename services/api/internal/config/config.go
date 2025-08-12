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
	MinioEndpointUrl   string
	MinioAccessId      string
	MinioSecretKey     string
	MinioStorageBucket string
	RabbitMQUri        string
	InputQueueName     string
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
		PostgresURI:        getEnv("POSTGRES_URI", "postgresql://root:password@localhost:5432/postgres"),
		MinioEndpointUrl:   getEnv("MINIO_ENDPOINT_URL", "localhost:9000"),
		MinioAccessId:      getEnv("MINIO_ACCESS_ID", "admin"),
		MinioSecretKey:     getEnv("MINIO_SECRET_KEY", "admin123"),
		MinioStorageBucket: getEnv("MINIO_STORAGE_BUCKET", "scripts"),
		RabbitMQUri:        getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672"),
		InputQueueName:     getEnv("INPUT_QUEUE_NAME", "input_queue"),
	}

	return config, err
}
