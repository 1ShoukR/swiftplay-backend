package models

import (
	"time"
)

type Message struct {
	MessageID int        `json:"message_id" db:"message_id"`
	MatchID   int        `json:"match_id" db:"match_id"`
	SenderID  int        `json:"sender_id" db:"sender_id"`
	Content   string     `json:"content" db:"content"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	ReadAt    *time.Time `json:"read_at,omitempty" db:"read_at"`
}
