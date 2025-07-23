package models

import (
	"time"
)

type Message struct {
	MessageID uint       `json:"message_id" gorm:"primaryKey;autoIncrement;column:message_id"`
	MatchID   uint       `json:"match_id" gorm:"not null;index;column:match_id"`
	SenderID  uint       `json:"sender_id" gorm:"not null;index;column:sender_id"`
	Content   string     `json:"content" gorm:"not null;type:text"`
	CreatedAt time.Time  `json:"created_at"`
	ReadAt    *time.Time `json:"read_at,omitempty" gorm:"column:read_at"`
}
