package handlers

import (
	"net/http"

	"github.com/1shoukr/swiftplay-backend/internal/database"
	"github.com/1shoukr/swiftplay-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *database.Database) *UserHandler {
	return &UserHandler{db: db.GetDB()}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var requestData struct {
		User    models.User    `json:"user"`
		Profile models.Profile `json:"profile"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}
	
	if requestData.User.Username == "" || requestData.User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and email are required",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.User.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
	}
	requestData.User.PasswordHash = string(hashedPassword)
	requestData.User.AuthLevel = "user"


	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&requestData.User).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": "Username or email might already exist",
		})
		return
	}

	requestData.Profile.UserID = requestData.User.UserID

	if requestData.Profile.GameRanks == nil {
		requestData.Profile.GameRanks = make(map[string]string)
	}

	if err := tx.Create(&requestData.Profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create profile",
			"details": err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to commit transaction",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":         "User and gaming profile created successfully",
		"user":            requestData.User,
		"profile":         requestData.Profile,
		"linked_games": []string{"valorant"}, // we can add the user's linked games here
	})
}
