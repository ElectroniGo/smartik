package config

import "os"

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
	return Config{
		Debug: getEnv("DEBUG", ""),
		Port:  getEnv("PORT", "8080"),
	}
}
