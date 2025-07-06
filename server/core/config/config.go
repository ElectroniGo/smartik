package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug string
	Port  string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err == nil {
		return Config{
			Debug: "true",
			Port:  os.Getenv("PORT"),
		}
	}

	return Config{
		Debug: getEnv("DEBUG", ""),
		Port:  getEnv("PORT", "8080"),
	}
}
