package handlers

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	return &AuthHandler{db: db.GetDB()}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login endpoint",
		"status":  "success",
	})
}