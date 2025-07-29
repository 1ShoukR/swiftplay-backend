package middleware

import (
	"net/http"
	"strings"

	"github.com/1shoukr/swiftplay-backend/internal/jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT tokens and injects user data into context
func AuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Expected: Bearer <token>",
			})
			c.Abort()
			return
		}

		token := bearerToken[1]

		claims, err := jwtService.ValidateAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid or expired token",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// RequireAuth is a basic auth middleware that just validates tokens
func RequireAuth(jwtService *jwt.JWTService) gin.HandlerFunc {
	return AuthMiddleware(jwtService)
}

// RequireUser ensures the user has at least "user" level access
func RequireUser(jwtService *jwt.JWTService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// First validate the token
		AuthMiddleware(jwtService)(c)
		if c.IsAborted() {
			return
		}

		// Check role (user is the minimum level)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User role not found in context",
			})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid role format",
			})
			c.Abort()
			return
		}

		validRoles := []string{"user", "admin", "super_admin", "engineer"}
		if !contains(validRoles, roleStr) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions. User access required.",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

// RequireAdmin ensures the user has admin level access or higher
func RequireAdmin(jwtService *jwt.JWTService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// First validate the token
		AuthMiddleware(jwtService)(c)
		if c.IsAborted() {
			return
		}

		// Check role
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User role not found in context",
			})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid role format",
			})
			c.Abort()
			return
		}

		validRoles := []string{"admin", "super_admin", "engineer"}
		if !contains(validRoles, roleStr) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions. Admin access required.",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

// RequireSuperAdmin ensures the user has super_admin level access or engineer
func RequireSuperAdmin(jwtService *jwt.JWTService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// First validate the token
		AuthMiddleware(jwtService)(c)
		if c.IsAborted() {
			return
		}

		// Check role
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User role not found in context",
			})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid role format",
			})
			c.Abort()
			return
		}

		validRoles := []string{"super_admin", "engineer"}
		if !contains(validRoles, roleStr) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions. Super admin access required.",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

// RequireEngineer ensures the user has engineer level access (highest level)
func RequireEngineer(jwtService *jwt.JWTService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// First validate the token
		AuthMiddleware(jwtService)(c)
		if c.IsAborted() {
			return
		}

		// Check role
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User role not found in context",
			})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid role format",
			})
			c.Abort()
			return
		}

		if roleStr != "engineer" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions. Engineer access required.",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// GetUserFromContext extracts user information from the Gin context
func GetUserFromContext(c *gin.Context) (userID uint, email string, role string, exists bool) {
	userIDVal, userIDExists := c.Get("user_id")
	emailVal, emailExists := c.Get("email")
	roleVal, roleExists := c.Get("role")

	if !userIDExists || !emailExists || !roleExists {
		return 0, "", "", false
	}

	userID, userIDOk := userIDVal.(uint)
	email, emailOk := emailVal.(string)
	role, roleOk := roleVal.(string)

	if !userIDOk || !emailOk || !roleOk {
		return 0, "", "", false
	}

	return userID, email, role, true
}
