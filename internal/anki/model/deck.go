package model

import "github.com/google/uuid"

type Deck struct {
	Id          uuid.UUID
	Topic       string
	Description string
	References  []string
}
