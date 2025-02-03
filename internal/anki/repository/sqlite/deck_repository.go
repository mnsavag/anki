package sqlite

import (
	"context"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/internal/anki/repository/dto"
)

func (r *SqliteRepository) AddDeck(ctx context.Context, deck *model.Deck) (uuid.UUID, error) {
	panic("implement me")
}
func (r *SqliteRepository) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
func (r *SqliteRepository) GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	panic("implement me")
}
func (r *SqliteRepository) UpdateDeck(ctx context.Context, id uuid.UUID, fieldsUpdate dto.UpdateDeckData) error {
	panic("implement me")
}
