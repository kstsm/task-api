package v1

import (
	"github.com/go-chi/chi/v5"
	"task-manager/internal/api/rest/v1/task"
	"task-manager/internal/service"
)

type Handler struct {
	service     service.ServiceI
	taskHandler *task.Handler
}

func NewHandler(service service.ServiceI) *Handler {
	return &Handler{
		service:     service,
		taskHandler: task.NewHandler(service),
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/task", func(r chi.Router) {
			h.taskHandler.SetupRoutes(r)
		})
	})
}
