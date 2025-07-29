package routes

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/handlers"
	"github.com/1shoukr/swiftplay-backend/internal/jwt"
	"github.com/1shoukr/swiftplay-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup, db *database.Database, jwtService *jwt.JWTService) {
	userHandler := handlers.NewUserHandler(db)

	users := api.Group("/users")
	{
		users.POST("/create", userHandler.CreateUser)

		users.GET("/profile", middleware.RequireUser(jwtService), getUserProfile)
		users.GET("/admin-only", middleware.RequireAdmin(jwtService), adminOnlyEndpoint)
		users.GET("/super-admin-only", middleware.RequireSuperAdmin(jwtService), superAdminOnlyEndpoint)
		users.GET("/engineer-only", middleware.RequireEngineer(jwtService), engineerOnlyEndpoint)
	}
}

// Example protected endpoint handlers to demonstrate JWT middleware
func getUserProfile(c *gin.Context) {
	userID, email, role, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User data not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User profile accessed successfully",
		"user": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
	})
}

func adminOnlyEndpoint(c *gin.Context) {
	userID, email, role, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User data not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin access granted",
		"user": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
		"data": "This is admin-only data",
	})
}

func superAdminOnlyEndpoint(c *gin.Context) {
	userID, email, role, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User data not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super admin access granted",
		"user": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
		"data": "This is super admin-only data",
	})
}

func engineerOnlyEndpoint(c *gin.Context) {
	userID, email, role, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User data not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Engineer access granted",
		"user": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
		"data": "This is engineer-only data",
	})
}
