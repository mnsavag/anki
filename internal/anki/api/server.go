package api

import (
	"context"

	"github.com/mnsavag/anki.git/internal/lib/log"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"
)

type Server struct {
	v1.UnimplementedAnkiServiceV1Server
	logger *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) CreateCard(ctx context.Context, in *v1.CreateCardRequest) (*v1.CreateCardResponse, error) {
	return &v1.CreateCardResponse{
		Id: "Hello world",
	}, nil
}
