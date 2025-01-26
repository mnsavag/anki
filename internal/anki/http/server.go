package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mnsavag/anki.git/internal/anki/config"
	"github.com/mnsavag/anki.git/internal/lib/log"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"
)

type Server struct {
	cfg        config.Config
	httpServer *http.Server
	logger     *log.Logger
}

func NewServer(cfg config.Config, logger *log.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Start() error {
	fmt.Printf("Server start on port: %s", s.cfg.HTTPPort)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Register(ctx context.Context) error {
	err := s.newHttpServer(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot register http server")
	}
	return nil
}

func (s *Server) newHttpServer(ctx context.Context) error {
	handler := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	s.httpServer = &http.Server{
		Addr:         net.JoinHostPort(s.cfg.Service.Host, s.cfg.Service.HTTPPort),
		Handler:      handler,
		IdleTimeout:  90 * time.Second,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	err := v1.RegisterAnkiServiceV1HandlerFromEndpoint(
		ctx,
		handler,
		net.JoinHostPort(s.cfg.Service.Host, s.cfg.Service.GRPCPort),
		opts,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	graceCtx, graceCancel := context.WithTimeout(context.Background(), s.cfg.Service.ShutdownContextTimeout)
	defer graceCancel()

	if err := s.httpServer.Shutdown(graceCtx); err != nil {
		return errors.Wrap(err, "could not gracefully shutdown http server")
	}

	// s.logger.Info("http servers stopped...")

	return nil
}
