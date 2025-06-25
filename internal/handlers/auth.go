package handlers

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/jmoiron/sqlx"
)

type AuthHandler struct {
	db *sqlx.DB
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	return &AuthHandler{db: db.GetDB()}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login endpoint"))
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create endpoint"))
}
