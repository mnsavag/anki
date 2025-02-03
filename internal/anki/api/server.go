package api

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/service"
	"github.com/mnsavag/anki.git/internal/lib/log"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) CreateDeck(ctx context.Context, in *v1.CreateDeckRequest) (*v1.CreateDeckResponse, error) {
	l := s.logger.WithContext(ctx)

	mappedRequest := fromCreateDeckRequest(in)

	id, err := s.ankiService.CreateDeck(ctx, mappedRequest)
	if err != nil {
		l.WithError(err).Error("CreateDeck request error")
		return nil, toGRPCError(err)
	}

	return &v1.CreateDeckResponse{
		Id: id,
	}, nil
}

func (s *Server) GetDeckById(ctx context.Context, in *v1.GetDeckByIdRequest) (*v1.GetDeckByIdResponse, error) {
	l := s.logger.WithContext(ctx)

	deckId, err := uuid.Parse(in.GetId())
	if err != nil {
		l.WithError(err).Error("invalid deck id")
		return nil, status.Error(codes.InvalidArgument, "invalid deck id")
	}

	deck, err := s.ankiService.GetDeckById(ctx, deckId)
	if err != nil {
		if errors.Is(err, service.ErrDeckNotFound) {
			return emptyGetDeckResponse(), nil
		}

		l.WithError(err).Error("GetDeckById request error")
		return nil, toGRPCError(err)
	}

	return toGetDeckResponse(deck), nil
}

func (s *Server) DeleteDeck(ctx context.Context, in *v1.DeleteDeckRequest) (*emptypb.Empty, error) {
	l := s.logger.WithContext(ctx)

	deckId, err := uuid.Parse(in.GetId())
	if err != nil {
		l.WithError(err).Error("invalid deck id")
		return nil, status.Error(codes.InvalidArgument, "invalid deck id")
	}

	err = s.ankiService.DeleteDeck(ctx, deckId)
	if err != nil {
		if errors.Is(err, service.ErrDeckNotFound) {
			return &emptypb.Empty{}, nil
		}

		l.WithError(err).Error("DeleteDeck request error")
		return nil, toGRPCError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateDeck(ctx context.Context, in *v1.UpdateDeckRequest) (*emptypb.Empty, error) {
	l := s.logger.WithContext(ctx)

	deckId, err := uuid.Parse(in.GetId())
	if err != nil {
		l.WithError(err).Error("invalid deck id")
		return nil, status.Error(codes.InvalidArgument, "invalid deck id")
	}

	updateDeckData := toUpdateDeckData(in)

	err = s.ankiService.UpdateDeck(ctx, deckId, updateDeckData)
	if err != nil {
		if errors.Is(err, service.ErrDeckNotFound) {
			return &emptypb.Empty{}, nil
		}

		l.WithError(err).Error("UpdateDeck request error")
		return nil, toGRPCError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) CreateCard(ctx context.Context, in *v1.CreateCardRequest) (*v1.CreateCardResponse, error) {
	l := s.logger.WithContext(ctx)

	deckId, err := uuid.Parse(in.GetDeckId())
	if err != nil {
		l.WithError(err).Error("invalid deck id")
		return nil, status.Error(codes.InvalidArgument, "invalid deck id")
	}

	createCardData := toCreateCardData(in)

	cardId, err := s.ankiService.CreateCard(ctx, deckId, createCardData)
	if err != nil {
		l.WithError(err).Error("CreateCard request error")
		return nil, toGRPCError(err)
	}

	return &v1.CreateCardResponse{
		Id: cardId,
	}, nil
}

func (s *Server) DeleteCard(ctx context.Context, in *v1.DeleteCardRequest) (*emptypb.Empty, error) {
	l := s.logger.WithContext(ctx)

	cardId, err := uuid.Parse(in.GetCardId())
	if err != nil {
		l.WithError(err).Error("invalid card id")
		return nil, status.Error(codes.InvalidArgument, "invalid card id")
	}

	err = s.ankiService.DeleteCard(ctx, cardId)
	if err != nil {
		if errors.Is(err, service.ErrCardNotFound) {
			return &emptypb.Empty{}, nil
		}

		l.WithError(err).Error("DeleteCard request error")
		return nil, toGRPCError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateCard(ctx context.Context, in *v1.UpdateCardRequest) (*emptypb.Empty, error) {
	l := s.logger.WithContext(ctx)

	cardId, err := uuid.Parse(in.GetCardId())
	if err != nil {
		l.WithError(err).Error("invalid card id")
		return nil, status.Error(codes.InvalidArgument, "invalid card id")
	}

	mappedRequest := toUpdateCardData(in)

	err = s.ankiService.UpdateCard(ctx, cardId, mappedRequest)
	if err != nil {
		if errors.Is(err, service.ErrCardNotFound) {
			return &emptypb.Empty{}, nil
		}

		l.WithError(err).Error("UpdateCard request error")
		return nil, toGRPCError(err)
	}

	return &emptypb.Empty{}, nil
}
