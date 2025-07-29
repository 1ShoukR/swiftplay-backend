package config

import (
	"fmt"
	"os"
	"time"
)

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret             string
	Expiry             time.Duration
	RefreshTokenSecret string
	RefreshTokenExpiry time.Duration
}

// LoadJWTConfig loads JWT configuration from environment variables
func LoadJWTConfig() (*JWTConfig, error) {
	secret := getEnv("JWT_SECRET", "")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is required")
	}

	if len(secret) < 32 {
		return nil, fmt.Errorf("JWT_SECRET must be at least 32 characters long for security")
	}

	refreshSecret := getEnv("REFRESH_TOKEN_SECRET", "")
	if refreshSecret == "" {
		return nil, fmt.Errorf("REFRESH_TOKEN_SECRET environment variable is required")
	}

	if len(refreshSecret) < 32 {
		return nil, fmt.Errorf("REFRESH_TOKEN_SECRET must be at least 32 characters long for security")
	}

	expiryStr := getEnv("JWT_EXPIRY", "24h")
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY format: %w", err)
	}

	refreshExpiryStr := getEnv("REFRESH_TOKEN_EXPIRY", "168h") // 7 days default
	refreshExpiry, err := time.ParseDuration(refreshExpiryStr)
	if err != nil {
		return nil, fmt.Errorf("invalid REFRESH_TOKEN_EXPIRY format: %w", err)
	}

	return &JWTConfig{
		Secret:             secret,
		Expiry:             expiry,
		RefreshTokenSecret: refreshSecret,
		RefreshTokenExpiry: refreshExpiry,
	}, nil
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
