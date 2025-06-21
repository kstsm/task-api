package utils

import (
	"encoding/json"
	"github.com/gookit/slog"
	"net/http"
	"task-manager/models"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Error when encoding JSON:", err)
		http.Error(w, "Ошибка при обработке ответа", http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.Error{Message: message})
}
