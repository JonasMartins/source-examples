package routes

import (
	"net/http"
	"project/src/files_upload/internal/handler"
)

type Router struct {
	H *handler.Handler
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/upload/single":
		ro.H.UploadSingleFile(w, r)
	case "/upload/multiple":
		ro.H.UploadMultipleFiles(w, r)
	}
}

func NewRouter(h *handler.Handler) http.Handler {
	return &Router{
		H: h,
	}
}
