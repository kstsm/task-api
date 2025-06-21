package storage

import (
	"context"
	"github.com/google/uuid"
	"task-manager/internal/apperrors"
	"task-manager/models"
	"time"
)

type TaskStatus string

const (
	Pending = "pending"
	Success = "success"
	Failed  = "failed"
)

func (s *Storage) CreateTask(ctx context.Context) (*models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &models.Task{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Status:    Pending,
	}

	s.storage[task.ID] = task

	return task, nil
}

func (s *Storage) SaveTask(ctx context.Context, task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.storage[task.ID]; !exists {
		return apperrors.ErrTaskNotFound
	}

	s.storage[task.ID] = task
	return nil
}

func (s *Storage) GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.storage[id]
	if !ok {
		return nil, apperrors.ErrTaskNotFound
	}
	return task, nil
}

func (s *Storage) DeleteTask(ctx context.Context, id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[id]; !ok {
		return apperrors.ErrTaskNotFound
	}
	delete(s.storage, id)
	return nil
}
