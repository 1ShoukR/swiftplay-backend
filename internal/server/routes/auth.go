package routes

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(db *database.Database) http.Handler {
	r := chi.NewRouter()
	authHandler := handlers.NewAuthHandler(db)

	r.Get("/login", authHandler.Login)
	r.Post("/create", authHandler.CreateUser)

	return r
}
