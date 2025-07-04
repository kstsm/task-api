package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"task-manager/internal/api/rest/v1/task"
	"task-manager/internal/service"
)

type Handler struct {
	service     service.ServiceI
	taskHandler *task.Handler
	handler     *chi.Mux
}

func NewHandler(service service.ServiceI) *Handler {
	h := &Handler{
		service:     service,
		taskHandler: task.NewHandler(service),
		handler:     chi.NewMux(),
	}
	h.setupHandlers()
	return h
}

func (h *Handler) setupHandlers() {
	h.handler.Route("/v1", func(r chi.Router) {
		r.Route("/task", func(r chi.Router) {
			h.taskHandler.SetupHandlers(r)
		})
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
