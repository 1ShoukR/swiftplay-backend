package models

import (
	"time"
)

type User struct {
	UserID       int        `json:"user_id" db:"user_id"`
	Username     string     `json:"username" db:"username"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"-" db:"password_hash"`
	SoftDelete   bool       `json:"soft_delete" db:"soft_delete"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
