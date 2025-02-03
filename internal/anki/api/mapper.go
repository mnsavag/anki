package api

import (
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/internal/anki/service/dto"
	"github.com/mnsavag/anki.git/pkg/api"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"
)

func fromCreateDeckRequest(req *v1.CreateDeckRequest) dto.CreateDeckData {
	return dto.CreateDeckData{
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

func toUpdateDeckData(req *v1.UpdateDeckRequest) dto.UpdateDeckData {
	return dto.UpdateDeckData{
		Topic:       req.Topic,
		Description: req.Description,
		References:  req.References,
	}
}

func emptyGetDeckResponse() *v1.GetDeckByIdResponse {
	return &v1.GetDeckByIdResponse{}
}

func toCreateCardData(req *v1.CreateCardRequest) dto.CreateCardData {
	return dto.CreateCardData{
		Question: req.Question,
		Answer:   req.Answer,
	}
}

func toUpdateCardData(req *v1.UpdateCardRequest) dto.UpdateCardData {
	return dto.UpdateCardData{
		Question: req.Question,
		Answer:   req.Answer,
	}
}
