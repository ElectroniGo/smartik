package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug  string
	Port   string
	PdfURL string // For development purposes, you can set a default PDF URL here.
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
			Debug:  "true",
			Port:   os.Getenv("PORT"),
			PdfURL: os.Getenv("PDF_URL"), // set in .env file
		}
	}

	return Config{
		Debug: getEnv("DEBUG", ""),
		Port:  getEnv("PORT", "8080"),
	}
}
