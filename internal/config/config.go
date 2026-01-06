package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPPort      string
	PostgresDSN   string
	UserAPIURL    string
	ProductAPIURL string
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
	userAPIURL := os.Getenv("USER_API_URL")
	if userAPIURL == "" {
		return nil, fmt.Errorf("USER_API_URL is required")
	}
	productAPIURL := os.Getenv("CATALOG_API_URL")
	if productAPIURL == "" {
		return nil, fmt.Errorf("CATALOG_API_URL is required")
	}
	return &Config{
		HTTPPort:      port,
		PostgresDSN:   postgresDSN,
		UserAPIURL:    userAPIURL,
		ProductAPIURL: productAPIURL,
	}, nil
}
