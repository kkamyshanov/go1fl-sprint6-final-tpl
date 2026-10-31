package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

func New(logger *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.Index)
	r.Post("/upload", handlers.Upload)

	return &Server{
		Logger: logger,
		HTTPServer: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
