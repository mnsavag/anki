package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mnsavag/anki.git/internal/anki/config"
)

type Server struct {
	cfg        config.Config
	httpServer *http.Server
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:           net.JoinHostPort(s.cfg.Service.Host, s.cfg.Service.HTTPPort),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server start on port: %s", s.cfg.HTTPPort)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
