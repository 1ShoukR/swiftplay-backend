package routes

import (
	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup, db *database.Database) {
	userHandler := handlers.NewUserHandler(db)

	users := api.Group("/users")
	{
		users.POST("/create", userHandler.CreateUser) 
	}
}
