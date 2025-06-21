package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"net/http"
	"task-manager/internal/apperrors"
	"task-manager/internal/utils"
)

func (h Handler) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.CreateTask(r.Context())
	if err != nil {
		slog.Errorf("Error creating task: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "Ошибка при создании задачи")
		return
	}

	slog.Infof("Task created: %s", task.ID)
	utils.WriteJSON(w, http.StatusCreated, task)
}

func (h Handler) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		slog.Warnf("Invalid UUID format: %s", taskID)
		utils.WriteError(w, http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	task, err := h.service.GetTask(r.Context(), parsedTaskID)
	if err != nil {
		if errors.Is(err, apperrors.ErrTaskNotFound) {
			slog.Infof("Task not found: %s", parsedTaskID)
			utils.WriteError(w, http.StatusNotFound, "Задача не найдена")
		} else {
			slog.Errorf("Error getting task %s: %v", parsedTaskID, err)
			utils.WriteError(w, http.StatusInternalServerError, "Ошибка при получении задачи")
		}
		return
	}

	slog.Infof("Task not found: %s", parsedTaskID)
	utils.WriteJSON(w, http.StatusOK, task)
}

func (h Handler) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	validateTaskID, err := uuid.Parse(taskID)
	if err != nil {
		slog.Warnf("Invalid UUID format: %s", taskID)
		utils.WriteError(w, http.StatusBadRequest, "Неверный формат UUID")
		return
	}

	err = h.service.DeleteTask(r.Context(), validateTaskID)
	if err != nil {
		if errors.Is(err, apperrors.ErrTaskNotFound) {
			slog.Infof("Invalid UUID format: %s", validateTaskID)
			utils.WriteError(w, http.StatusNotFound, "Задача не найдена")
		} else {
			slog.Errorf("Error deleting task %s: %v", validateTaskID, err)
			utils.WriteError(w, http.StatusInternalServerError, "Ошибка при удалении задачи")
		}
		return
	}

	slog.Infof("Task deleted: %s", validateTaskID)
	utils.WriteJSON(w, http.StatusOK, "Задача удалена")
}
