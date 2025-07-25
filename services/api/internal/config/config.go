package config

import "os"

type Env struct {
	Port        string
	PostgresURI string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func Load() *Env {
	return &Env{
		Port:        getEnv("PORT", "1323"),
		PostgresURI: getEnv("POSTGRES_URI", "postgresql://root:123456@localhost:5432/postgres"),
	}
}
