package api

import (
	"context"
	"testing"

	"github.com/mnsavag/anki.git/internal/anki/api/mock"
	"github.com/mnsavag/anki.git/internal/anki/model"
	"github.com/mnsavag/anki.git/internal/anki/service"
	"github.com/mnsavag/anki.git/internal/lib/log"
	apiModel "github.com/mnsavag/anki.git/pkg/api"
	v1 "github.com/mnsavag/anki.git/pkg/api/v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

type ServerSuite struct {
	suite.Suite
	server      *Server
	ankiService *mock.MockAnkiService
	ctx         context.Context
}

func (s *ServerSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	l, _ := log.NewLogger(log.NewDefaultConfig())

	s.ankiService = mock.NewMockAnkiService(ctrl)
	s.server = NewServer(
		l,
		s.ankiService,
	)
	s.ctx = context.Background()
}

func (s *ServerSuite) TestGetDeckById() {
	//arrange
	deckId := uuid.NewString()
	deckUUID, _ := uuid.Parse(deckId)
	want := &v1.GetDeckByIdResponse{
		Deck: &apiModel.Deck{
			Topic:       "topic",
			Description: "desc",
			References:  []string{"ref1", "ref2"},
		},
	}

	//act
	s.ankiService.EXPECT().GetDeckById(gomock.Any(), deckUUID).Return(model.Deck{
		Topic:       "topic",
		Description: "desc",
		References:  []string{"ref1", "ref2"},
	}, nil)
	got, err := s.server.GetDeckById(
		s.ctx,
		&v1.GetDeckByIdRequest{
			Id: deckId,
		},
	)

	//assert
	s.Nil(err)
	s.Equal(got, want)
}

func (s *ServerSuite) TestGetDeckById_NotFound() {
	//arrange
	deckId := uuid.NewString()
	deckUUID, _ := uuid.Parse(deckId)
	want := &v1.GetDeckByIdResponse{}

	//act
	s.ankiService.EXPECT().GetDeckById(gomock.Any(), deckUUID).Return(model.Deck{}, service.ErrDeckNotFound)
	got, err := s.server.GetDeckById(
		s.ctx,
		&v1.GetDeckByIdRequest{
			Id: deckId,
		},
	)

	//assert
	s.ErrorIs(err, nil)
	s.Equal(got, want)
}
