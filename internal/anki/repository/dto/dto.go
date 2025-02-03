package dto

type UpdateDeckData struct {
	Topic       string
	Description string
	References  []string
}

type UpdateCardData struct {
	Question string
	Answer   string
}
