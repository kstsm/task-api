package task

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"net/http"
	"task-manager/internal/apperrors"
	"task-manager/internal/service"
	"task-manager/internal/utils"
)

type Handler struct {
	service service.ServiceI
}

func NewHandler(service service.ServiceI) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRoutes(r chi.Router) {
	r.Post("/", h.createTaskHandler)
	r.Get("/{id}", h.getTaskHandler)
	r.Delete("/{id}", h.deleteTaskHandler)
}

func (h *Handler) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warnf("Error parsing request: %v", err)
		utils.WriteError(w, http.StatusBadRequest, "Недопустимый запрос")
		return
	}

	task, err := h.service.CreateTask(r.Context())
	if err != nil {
		slog.Errorf("Error creating task: %v", err)
		utils.WriteError(w, http.StatusBadRequest, "Недопустимый запрос")
		return
	}

	response := CreateTaskResponse{
		ID:        task.ID,
		CreatedAt: task.CreatedAt,
	}

	slog.Infof("Task created: %s", task.ID)
	utils.WriteJSON(w, http.StatusCreated, response)
}

func (h *Handler) getTaskHandler(w http.ResponseWriter, r *http.Request) {
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
			utils.WriteError(w, http.StatusBadRequest, "Задача не найдена")
		} else {
			slog.Errorf("Error getting task %s: %v", parsedTaskID, err)
			errorResp := ErrorResponse{
				Error: "internal error",
			}
			utils.WriteJSON(w, http.StatusInternalServerError, errorResp)
		}

		return
	}

	response := GetTaskResponse{
		ID:        task.ID,
		Result:    task.Result,
		Duration:  task.Duration,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	}

	slog.Infof("Task not found: %s", parsedTaskID)
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
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
			slog.Infof("Task not found: %s", validateTaskID)
			utils.WriteError(w, http.StatusNotFound, "Задача не найдена")
		} else {
			slog.Errorf("Error deleting task %s: %v", validateTaskID, err)
			errorResp := ErrorResponse{
				Error: "internal error",
			}
			utils.WriteJSON(w, http.StatusInternalServerError, errorResp)
		}
		return
	}

	slog.Infof("Task deleted: %s", validateTaskID)
	utils.WriteJSON(w, http.StatusOK, nil)
}
