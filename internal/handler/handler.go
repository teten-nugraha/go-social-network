package handler

import (
	"github.com/matryer/way"
	"github.com/teten-nugraha/go-social-network/internal/service"
	"net/http"
)

type handler struct {
	service *service.Service
}

// New creates an http.Handler with predefined routing
func New(s *service.Service) http.Handler {

	h := &handler{s}

	api := way.NewRouter()
	api.HandleFunc("POST", "/login", h.login)
	api.HandleFunc("POST", "/users", h.createUser)

	r := way.NewRouter()
	r.Handle("*","/api...", http.StripPrefix("/api", api))

	return r
}