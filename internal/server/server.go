package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {

	r := chi.NewRouter()

	// Handlers
	// ...
	//

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}
}
