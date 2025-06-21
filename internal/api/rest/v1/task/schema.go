package task

import (
	"github.com/google/uuid"
	"time"
)

type CreateTaskRequest struct {
	Data int `json:"data"`
}
type CreateTaskResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTaskResponse struct {
	ID        uuid.UUID `json:"id"`
	Result    int       `json:"result"`
	Duration  int       `json:"duration"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
