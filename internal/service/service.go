package service

import (
	"context"
	"github.com/google/uuid"
	"task-manager/internal/storage"
	"task-manager/models"
)

type ServiceI interface {
	CreateTask(ctx context.Context) (*models.Task, error)
	processTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, taskID uuid.UUID) (*models.Task, error)
	DeleteTask(ctx context.Context, taskID uuid.UUID) error
}

type Service struct {
	storage storage.StorageI
}

func NewService(storage storage.StorageI) *Service {
	return &Service{
		storage: storage,
	}
}
