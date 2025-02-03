// mockgen -destination=./mock/repository_mock_generated.go -package=mock -source repository.go
package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mnsavag/anki.git/internal/anki/service/dto"
	"github.com/mnsavag/anki.git/internal/anki/service/mock"
	"github.com/mnsavag/anki.git/internal/lib/log"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite
	service *Service
	repo    *mock.MockRepository
	ctx     context.Context
}

func (s *ServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	l, _ := log.NewLogger(log.NewDefaultConfig())

	s.repo = mock.NewMockRepository(ctrl)
	s.service = NewService(
		l,
		s.repo,
	)
	s.ctx = context.Background()
}

func (s *ServiceSuite) TestCreateDeck() {
	//arrange
	id := uuid.New()
	want := id.String()

	s.repo.EXPECT().AddDeck(gomock.Any(), gomock.Any()).Return(id, nil)

	//act
	got, err := s.service.CreateDeck(s.ctx, dto.CreateDeckData{})

	//assert
	s.Nil(err)
	s.Equal(got, want)
}
