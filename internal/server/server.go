package server

import (
	"log"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	engine *gin.Engine
	DB     *database.Database
}

func NewServer() (*Server, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Set Gin mode to release in production
	gin.SetMode(gin.DebugMode)

	dbConfig := database.LoadConfig()
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()

	// Setup routes
	routes.SetupRoutes(engine, db)

	server := &Server{
		engine: engine,
		DB:     db,
	}

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
