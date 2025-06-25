package models

import (
	"time"
)

type Match struct {
	MatchID   int       `json:"match_id" db:"match_id"`
	UserID1   int       `json:"user_id_1" db:"user_id_1"`
	UserID2   int       `json:"user_id_2" db:"user_id_2"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
