package model

import "github.com/google/uuid"

type Card struct {
	Id       uuid.UUID
	Question string
	Answer   string
}
