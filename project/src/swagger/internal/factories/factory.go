package factories

import (
	"net/http"
	"project/pkg/utils"
	"project/src/swagger/internal/handler"
	"project/src/swagger/internal/handler/middleware"
	"project/src/swagger/internal/routes"
)

func BuildSrv(mux *http.ServeMux) *http.Server {
	h := newHandler()
	routes.AppRouter(mux, h)
	return &http.Server{
		Addr:    utils.DEFAULT_PORT_STR,
		Handler: middleware.Log(mux),
	}
}

func newHandler() *handler.Handler {
	return &handler.Handler{}
}
