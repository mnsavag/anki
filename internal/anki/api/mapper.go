package api

import (
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/pkg/api"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"
)

func fromCreateDeckRequest(req *v1.CreateDeckRequest) model.CreateDeckRequest {
	return model.CreateDeckRequest{
		Topic:       req.Topic,
		Description: req.Description,
		References:  req.References,
	}
}

func toGetDeckResponse(deck model.Deck) *v1.GetDeckByIdResponse {
	return &v1.GetDeckByIdResponse{
		Deck: &api.Deck{
			Topic:       deck.Topic,
			Description: deck.Description,
			References:  deck.References,
		},
	}
}

func emptyGetDeckResponse() *v1.GetDeckByIdResponse {
	return &v1.GetDeckByIdResponse{}
}
