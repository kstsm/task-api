package storage

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"task-manager/models"
)

type StorageI interface {
	CreateTask(ctx context.Context) (*models.Task, error)
	SaveTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, taskID uuid.UUID) (*models.Task, error)
	DeleteTask(ctx context.Context, taskID uuid.UUID) error
}

type Storage struct {
	mu      sync.RWMutex
	storage map[uuid.UUID]*models.Task
}

func NewStorage() StorageI {
	return &Storage{
		storage: make(map[uuid.UUID]*models.Task),
	}
}
