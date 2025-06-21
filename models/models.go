package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Result    int       `json:"result"`
	Duration  int       `json:"duration"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Error struct {
	Message string `json:"message"`
}
