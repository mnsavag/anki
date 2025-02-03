package repository

import "errors"

var (
	ErrDeckNotFound = errors.New("deck not found")
	ErrCardNotFound = errors.New("card not found")
)
