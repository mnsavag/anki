package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/mnsavag/anki.git/internal/anki/config"
	"github.com/mnsavag/anki.git/internal/lib/log"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"
	"github.com/pkg/errors"
)

type Server struct {
	server *grpc.Server
	cfg    config.Config
	logger *log.Logger
}

func NewServer(cfg config.Config, logger *log.Logger) *Server {
	return &Server{
		server: grpc.NewServer(),
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) RegisterAnkiV1(apiServer v1.AnkiServiceV1Server) {
	v1.RegisterAnkiServiceV1Server(s.server, apiServer)
}

func (s *Server) Start() error {
	listener, err := net.Listen(
		"tcp",
		net.JoinHostPort(s.cfg.Service.Host, s.cfg.Service.GRPCPort),
	)
	if err != nil {
		return errors.Wrap(err, "failed to listen gRPC server")
	}

	go func() {
		if err := s.server.Serve(listener); err != nil {
			// TODO: log.Fatal("gRPC server start failed")
			fmt.Println("gRPC server start failed")
		}
	}()
	// TODO: log.Fatal("gRPC server started, port: %s")
	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
	//s.logger.Info("gRPC server stopped...")
}
