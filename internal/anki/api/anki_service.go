// mockgen -destination=./mock/anki_service_mock_generated.go -package=mock -source anki_service.go
package api

import (
	"context"

	"github.com/mnsavag/anki.git/internal/anki/model"

	"github.com/google/uuid"
)

type AnkiService interface {
	CardExecutor
	DeckExecutor
}

type CardExecutor interface {
	CreateCard(ctx context.Context, deckId uuid.UUID) (string, error)
	DeleteCard(ctx context.Context, cardId uuid.UUID) error
	UpdateCard(ctx context.Context, cardId uuid.UUID) error
}

type DeckExecutor interface {
	CreateDeck(ctx context.Context, request model.CreateDeckRequest) (string, error)
	DeleteDeck(ctx context.Context, id uuid.UUID) error
	GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error)
	UpdateDeck(ctx context.Context, id uuid.UUID) error
}
