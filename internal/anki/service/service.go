package service

import (
	"context"
	"errors"

	"github.com/mnsavag/anki.git/internal/anki/model"
	repo "github.com/mnsavag/anki.git/internal/anki/repository"
	repoDto "github.com/mnsavag/anki.git/internal/anki/repository/dto"
	"github.com/mnsavag/anki.git/internal/anki/service/dto"
	"github.com/mnsavag/anki.git/internal/lib/log"

	"github.com/google/uuid"
)

type Service struct {
	logger *log.Logger
	repo   Repository
}

func NewService(logger *log.Logger, repo Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) CreateDeck(ctx context.Context, createData dto.CreateDeckData) (string, error) {
	l := s.logger.WithContext(ctx)

	deck := model.Deck{
		Id:          uuid.New(),
		Topic:       createData.Topic,
		Description: createData.Description,
		References:  createData.References,
	}

	id, err := s.repo.AddDeck(ctx, &deck)
	if err != nil {
		l.WithError(err).Error("cant create deck")
		return "", err
	}

	return id.String(), nil
}

func (s *Service) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	l := s.logger.WithContext(ctx)

	err := s.repo.DeleteDeck(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrDeckNotFound) {
			return ErrDeckNotFound
		}

		l.WithError(err).Error("delete deck error")
		return err
	}

	return nil
}

func (s *Service) GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	l := s.logger.WithContext(ctx)

	deck, err := s.repo.GetDeckById(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrDeckNotFound) {
			return model.Deck{}, ErrDeckNotFound
		}

		l.WithError(err).Error("get deck error")
		return model.Deck{}, err
	}

	return deck, nil
}

func (s *Service) UpdateDeck(ctx context.Context, id uuid.UUID, updateData dto.UpdateDeckData) error {
	// проверяем что колода существует, затем зная id в репо будем менять данные запросом
	l := s.logger.WithContext(ctx)

	_, err := s.GetDeckById(ctx, id)
	if err != nil {
		return err
	}

	err = s.repo.UpdateDeck(ctx, id, repoDto.UpdateDeckData{
		Topic:       updateData.Topic,
		Description: updateData.Description,
		References:  updateData.References,
	})
	if err != nil {
		l.WithError(err).Error("update deck error")
		return err
	}

	return nil
}

func (s *Service) CreateCard(ctx context.Context, deckId uuid.UUID, card dto.CreateCardData) (string, error) {
	l := s.logger.WithContext(ctx)

	_, err := s.GetDeckById(ctx, deckId)
	if err != nil {
		return "", err
	}

	id, err := s.repo.AddCard(ctx, deckId, &model.Card{
		Question: card.Question,
		Answer:   card.Answer,
	})
	if err != nil {
		l.WithError(err).Error("create card error")
		return "", err
	}

	return id.String(), nil
}

func (s *Service) DeleteCard(ctx context.Context, cardId uuid.UUID) error {
	l := s.logger.WithContext(ctx)

	err := s.repo.DeleteCard(ctx, cardId)
	if err != nil {
		if errors.Is(err, repo.ErrCardNotFound) {
			return ErrCardNotFound
		}

		l.WithError(err).Error("delete card error")
		return err
	}

	return nil
}

func (s *Service) UpdateCard(ctx context.Context, cardId uuid.UUID, card dto.UpdateCardData) error {
	l := s.logger.WithContext(ctx)

	err := s.repo.UpdateCard(ctx, cardId, repoDto.UpdateCardData{
		Question: card.Question,
		Answer:   card.Answer,
	})
	if err != nil {
		l.WithError(err).Error("update card error")
		return err
	}

	return nil
}
