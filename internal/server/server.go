package server

import (
	"log"
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/server/routes"
	"github.com/joho/godotenv"
)

type Server struct {
	*http.Server
	DB *database.Database
}

func NewServer() (*Server, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Initialize database connection
	dbConfig := database.LoadConfig()
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		return nil, err
	}

	// Create router with database dependency
	router := routes.NewRouter(db)

	server := &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		DB: db,
	}

	return server, nil
}

func (s *Server) Close() error {
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}
