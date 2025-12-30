package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPPort    string
	PostgresDSN string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		return nil, fmt.Errorf("POSTGRES_DSN is required")
	}
	return &Config{
		HTTPPort:    port,
		PostgresDSN: postgresDSN,
	}, nil
}
