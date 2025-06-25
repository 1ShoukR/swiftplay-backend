package routes

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *database.Database) http.Handler {
	r := chi.NewRouter()

	// Logger middleware
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Mount auth routes under /api/auth
	r.Mount("/api/auth", NewAuthRouter(db))

	return r
}
