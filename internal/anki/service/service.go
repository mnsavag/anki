package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/internal/lib/log"
)

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) CreateDeck(ctx context.Context, request model.CreateDeckRequest) (string, error) {
	panic("implement me")
}

func (s *Service) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (s *Service) GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	panic("implement me")
}

func (s *Service) UpdateDeck(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (s *Service) CreateCard(ctx context.Context, deckId uuid.UUID) (string, error) {
	panic("implement me")
}

func (s *Service) DeleteCard(ctx context.Context, cardId uuid.UUID) error {
	panic("implement me")
}

func (s *Service) UpdateCard(ctx context.Context, cardId uuid.UUID) error {
	panic("implement me")
}
