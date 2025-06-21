package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"math/rand"
	"task-manager/internal/storage"
	"task-manager/models"
	"time"
)

func (s Service) CreateTask(ctx context.Context) (*models.Task, error) {
	task, err := s.storage.CreateTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.storage.CreateTask: %w", err)
	}

	go func(t *models.Task) {
		err = s.processTask(context.Background(), t)
		if err != nil {
			slog.Errorf("task %v failed with error: %v", task.ID, err)
		}
	}(task)

	return task, nil
}

func (s Service) processTask(ctx context.Context, task *models.Task) error {
	start := time.Now()
	slog.Infof("Task %s is running", task.ID)

	// Имитация работы долгой задачи
	select {
	case <-time.After(time.Duration(10+rand.Intn(30)) * time.Second):
		task.Status = storage.Success
	case <-ctx.Done():
		task.Status = storage.Failed
	}

	task.Duration = int(time.Since(start).Seconds())
	task.Result = task.Duration * task.Duration

	err := s.storage.SaveTask(ctx, task)
	if err != nil {
		slog.Errorf("Error saving the task result %s: %v", task.ID, err)
		return fmt.Errorf("s.storage.SaveTask %w", err)
	}

	slog.Infof("Task %s completed with the status: %s", task.ID, task.Status)
	return nil
}

func (s Service) GetTask(ctx context.Context, taskID uuid.UUID) (*models.Task, error) {
	return s.storage.GetTask(ctx, taskID)
}

func (s Service) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	return s.storage.DeleteTask(ctx, taskID)
}
