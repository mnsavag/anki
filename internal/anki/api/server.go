package api

import (
	"context"

	"github.com/mnsavag/anki.git/internal/lib/log"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	v1.UnimplementedAnkiServiceV1Server
	logger      *log.Logger
	ankiService AnkiService
}

func NewServer(logger *log.Logger, ankiService AnkiService) *Server {
	return &Server{
		logger:      logger,
		ankiService: ankiService,
	}
}

func (s *Server) CreateCard(ctx context.Context, in *v1.CreateCardRequest) (*v1.CreateCardResponse, error) {
	return &v1.CreateCardResponse{
		Id: "Hello world",
	}, nil
}

func (s *Server) CreateDeck(ctx context.Context, in *v1.CreateDeckRequest) (*v1.CreateDeckResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteCard(ctx context.Context, in *v1.DeleteCardRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *Server) DeleteDeck(ctx context.Context, in *v1.DeleteDeckRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *Server) GetDeckById(ctx context.Context, in *v1.GetDeckByIdRequest) (*v1.GetDeckByIdResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateCard(ctx context.Context, in *v1.UpdateCardRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *Server) UpdateDeck(ctx context.Context, in *v1.UpdateDeckRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
