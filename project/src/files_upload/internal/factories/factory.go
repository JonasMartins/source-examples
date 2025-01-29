package factories

import (
	"net/http"
	"project/pkg/utils"
	"project/src/files_upload/internal/handler"
	"project/src/files_upload/internal/routes"
)

func BuildSrv() *http.Server {
	return &http.Server{
		Addr: utils.DEFAULT_PORT_STR,
		Handler: &routes.Router{
			H: &handler.Handler{},
		},
	}
}
