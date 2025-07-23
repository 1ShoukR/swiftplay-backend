package routes

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *database.Database) {
	// Logger middleware (Gin has built-in logger)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// API route group
	api := r.Group("/api")
	{
		// Mount auth routes under /api/auth
		SetupAuthRoutes(api, db)
		
		// Mount user routes under /api/users
		SetupUserRoutes(api, db)
	}
}
