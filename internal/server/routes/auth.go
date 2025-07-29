package routes

import (
	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/handlers"
	"github.com/1shoukr/swiftplay-backend/internal/jwt"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(api *gin.RouterGroup, db *database.Database, jwtService *jwt.JWTService) {
	authHandler := handlers.NewAuthHandler(db)

	auth := api.Group("/auth")
	{
		auth.GET("/login", authHandler.Login)
	}
}
