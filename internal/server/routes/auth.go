package routes

import (
	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(api *gin.RouterGroup, db *database.Database) {
	authHandler := handlers.NewAuthHandler(db)

	auth := api.Group("/auth")
	{
		auth.GET("/login", authHandler.Login)
	}
}
