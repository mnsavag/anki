package api

import "context"

type AnkiService interface {
	CardExecutor
	DeckExecutor
}

type CardExecutor interface {
	CreateCard(ctx context.Context)
	DeleteCard(ctx context.Context)
	UpdateCard(ctx context.Context)
}

type DeckExecutor interface {
	CreateDeck(ctx context.Context)
	DeleteDeck(ctx context.Context)
	GetDeckById(ctx context.Context)
	UpdateDeck(ctx context.Context)
}
