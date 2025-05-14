package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)

	if ok {
		return v
	}

	return fallback
}
