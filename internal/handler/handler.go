package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	v1 "task-manager/internal/api/rest/v1"
	"task-manager/internal/service"
)

type HandlerI interface {
	NewRouter() http.Handler
}

type Handler struct {
	service service.ServiceI
}

func NewHandler(svc service.ServiceI) HandlerI {
	return &Handler{
		service: svc,
	}
}

func (h Handler) NewRouter() http.Handler {
	r := chi.NewRouter()

	v1Handler := v1.NewHandler(h.service)
	v1Handler.SetupRoutes(r)

	return r
}
