package handlers

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
)

type AuthHandler struct {
	db *database.Database
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login endpoint"))
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create endpoint"))
}
