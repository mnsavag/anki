package dto

type CreateDeckData struct {
	Topic       string
	Description string
	References  []string
}

type CreateCardData struct {
	Question string
	Answer   string
}

type UpdateDeckData struct {
	Topic       string
	Description string
	References  []string
}

type UpdateCardData struct {
	Question string
	Answer   string
}
