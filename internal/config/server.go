package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/1shoukr/swiftplay-backend/internal/database"
)

// ServerConfig holds all server configuration
type ServerConfig struct {
	Port    int
	GinMode string
	DB      *database.Config
	JWT     *JWTConfig
}

// LoadServerConfig loads all configuration from environment variables
func LoadServerConfig() (*ServerConfig, error) {
	jwtConfig, err := LoadJWTConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load JWT configuration: %w", err)
	}

	dbConfig := database.LoadConfig()

	portStr := getEnv("PORT", "8081")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT value: %w", err)
	}

	ginMode := getEnv("GIN_MODE", "debug")
	if ginMode != "debug" && ginMode != "release" && ginMode != "test" {
		return nil, fmt.Errorf("invalid GIN_MODE value: %s (must be debug, release, or test)", ginMode)
	}

	return &ServerConfig{
		Port:    port,
		GinMode: ginMode,
		DB:      dbConfig,
		JWT:     jwtConfig,
	}, nil
}

// ValidateRequiredEnvVars validates that all required environment variables are set
func ValidateRequiredEnvVars() error {
	requiredVars := []string{
		"JWT_SECRET",
		"REFRESH_TOKEN_SECRET",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
	}

	var missingVars []string
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			missingVars = append(missingVars, varName)
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return nil
}
