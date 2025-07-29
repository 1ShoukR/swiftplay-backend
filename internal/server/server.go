package server

import (
	"fmt"
	"log"

	"github.com/1shoukr/swiftplay-backend/internal/config"
	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/jwt"
	"github.com/1shoukr/swiftplay-backend/internal/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	engine     *gin.Engine
	DB         *database.Database
	Config     *config.ServerConfig
	JWTService *jwt.JWTService
}

func NewServer() (*Server, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	if err := config.ValidateRequiredEnvVars(); err != nil {
		return nil, fmt.Errorf("environment validation failed: %w", err)
	}

	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load server configuration: %w", err)
	}

	gin.SetMode(serverConfig.GinMode)

	db, err := database.NewDatabase(serverConfig.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	jwtService := jwt.NewJWTService(serverConfig.JWT)

	engine := gin.Default()

	routes.SetupRoutes(engine, db, jwtService)

	server := &Server{
		engine:     engine,
		DB:         db,
		Config:     serverConfig,
		JWTService: jwtService,
	}

	log.Printf("Server configured successfully - Port: %d, Gin Mode: %s",
		serverConfig.Port, serverConfig.GinMode)
	log.Printf("JWT configuration loaded - Token expiry: %v, Refresh expiry: %v",
		serverConfig.JWT.Expiry, serverConfig.JWT.RefreshTokenExpiry)

	return server, nil
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) Close() error {
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}
