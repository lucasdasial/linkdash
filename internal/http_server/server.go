package httpserver

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	mux := http.NewServeMux()

	registerRoutes(mux)

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: loggingMiddleware(mux),
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server running on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
