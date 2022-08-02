package handler

import (
	"html/template"

	"github.com/gorilla/mux"
	"github.com/khusainnov/task3/pkg/service"
)

var templates = template.Must(template.ParseFiles("index.html"))

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/upload", h.UploadSystem)

	return r
}
