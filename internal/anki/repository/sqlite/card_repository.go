package sqlite

import (
	"context"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/internal/anki/repository/dto"
)

func (r *SqliteRepository) AddCard(ctx context.Context, deckId uuid.UUID, card *model.Card) (uuid.UUID, error) {
	panic("implement me")
}
func (r *SqliteRepository) DeleteCard(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
func (r *SqliteRepository) UpdateCard(ctx context.Context, id uuid.UUID, card dto.UpdateCardData) error {
	panic("implement me")
}
