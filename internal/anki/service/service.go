package service

import (
	"context"

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

func (s *Service) CreateCard(ctx context.Context) {
	panic("implement me")
}

func (s *Service) CreateDeck(ctx context.Context) {
	panic("implement me")
}

func (s *Service) DeleteCard(ctx context.Context) {
	panic("implement me")
}

func (s *Service) DeleteDeck(ctx context.Context) {
	panic("implement me")
}

func (s *Service) GetDeckById(ctx context.Context) {
	panic("implement me")
}

func (s *Service) UpdateCard(ctx context.Context) {
	panic("implement me")
}

func (s *Service) UpdateDeck(ctx context.Context) {
	panic("implement me")
}
