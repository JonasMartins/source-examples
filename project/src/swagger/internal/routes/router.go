package routes

import (
	"net/http"
	"project/src/swagger/internal/handler"
)

func AppRouter(mux *http.ServeMux, h *handler.Handler) {
	mux.HandleFunc("GET /health", h.Health)
}
