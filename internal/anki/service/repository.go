package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/model"
	repoDto "github.com/mnsavag/anki.git/internal/anki/repository/dto"
)

type Repository interface {
	DeckRepository
	CardRepository
}

type DeckRepository interface {
	AddDeck(ctx context.Context, deck *model.Deck) (uuid.UUID, error)
	DeleteDeck(ctx context.Context, id uuid.UUID) error
	GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error)
	UpdateDeck(ctx context.Context, id uuid.UUID, fieldsUpdate repoDto.UpdateDeckData) error
}

type CardRepository interface {
	AddCard(ctx context.Context, deckId uuid.UUID, card *model.Card) (uuid.UUID, error)
	DeleteCard(ctx context.Context, id uuid.UUID) error
	UpdateCard(ctx context.Context, id uuid.UUID, card repoDto.UpdateCardData) error
}
