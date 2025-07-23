package models

import (
	"time"
)

type Match struct {
	MatchID   uint      `json:"match_id" gorm:"primaryKey;autoIncrement;column:match_id"`
	UserID1   uint      `json:"user_id_1" gorm:"not null;index;column:user_id_1"`
	UserID2   uint      `json:"user_id_2" gorm:"not null;index;column:user_id_2"`
	Status    string    `json:"status" gorm:"not null;size:20;default:'pending'"`
	CreatedAt time.Time `json:"created_at"`
}
