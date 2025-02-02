package model

type CreateDeckRequest struct {
	Topic       string
	Description string
	References  []string
}

type Deck struct {
	Id          string
	Topic       string
	Description string
	References  []string
}
