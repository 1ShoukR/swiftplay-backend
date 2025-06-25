package models

import (
	"time"
)

type Profile struct {
	ProfileID    int        `json:"profile_id" db:"profile_id"`
	UserID       int        `json:"user_id" db:"user_id"`
	FirstName    *string    `json:"first_name,omitempty" db:"first_name"`
	LastName     *string    `json:"last_name,omitempty" db:"last_name"`
	Gender       *string    `json:"gender,omitempty" db:"gender"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Bio          *string    `json:"bio,omitempty" db:"bio"`
	City         *string    `json:"city,omitempty" db:"city"`
	Country      *string    `json:"country,omitempty" db:"country"`
	ValorantRank *string    `json:"valorant_rank,omitempty" db:"valorant_rank"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
