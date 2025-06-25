package routes

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(db *database.Database) http.Handler {
	r := chi.NewRouter()

	// Auth routes will go here
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		// You can now use the database connection here
		// Example: db.GetConn().Query(...)
		w.Write([]byte("Login endpoint"))
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		// You can now use the database connection here
		// Example: db.GetConn().Exec(...)
		w.Write([]byte("Register endpoint"))
	})

	return r
}
