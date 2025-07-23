package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       uint           `json:"user_id" gorm:"primaryKey;autoIncrement;column:user_id"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
	PasswordHash string         `json:"-" gorm:"not null;column:password_hash"`
	SoftDelete   bool           `json:"soft_delete" gorm:"default:false;column:soft_delete"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}


type Profile struct {
	ProfileID    uint       `json:"profile_id" gorm:"primaryKey;autoIncrement;column:profile_id"`
	UserID       uint       `json:"user_id" gorm:"not null;index;column:user_id"`
	FirstName    *string    `json:"first_name,omitempty" gorm:"size:50;column:first_name"`
	LastName     *string    `json:"last_name,omitempty" gorm:"size:50;column:last_name"`
	Gender       *string    `json:"gender,omitempty" gorm:"size:20"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty" gorm:"column:date_of_birth"`
	Bio          *string    `json:"bio,omitempty" gorm:"type:text"`
	City         *string    `json:"city,omitempty" gorm:"size:100"`
	Country      *string    `json:"country,omitempty" gorm:"size:100"`
	ValorantRank *string    `json:"valorant_rank,omitempty" gorm:"size:50;column:valorant_rank"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
