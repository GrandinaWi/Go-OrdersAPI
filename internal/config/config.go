package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPPort      string
	JWTSecret     []byte
	PostgresDSN   string
	UserAPIURL    string
	ProductAPIURL string
	OrdersAPIURL  string
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
	productAPIURL := os.Getenv("PRODUCT_API_URL")
	if productAPIURL == "" {
		return nil, fmt.Errorf("PRODUCT_API_URL is required")
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	ordersAPIURL := os.Getenv("ORDERS_API_URL")
	if ordersAPIURL == "" {
		return nil, fmt.Errorf("ORDERS_API_URL is required")
	}
	return &Config{
		HTTPPort:      port,
		JWTSecret:     jwtSecret,
		PostgresDSN:   postgresDSN,
		UserAPIURL:    userAPIURL,
		ProductAPIURL: productAPIURL,
		OrdersAPIURL:  ordersAPIURL,
	}, nil
}
