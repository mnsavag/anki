// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -destination=./mock/repository_mock_generated.go -package=mock -source repository.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	model "github.com/mnsavag/anki.git/internal/anki/model"
	dto "github.com/mnsavag/anki.git/internal/anki/repository/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddCard mocks base method.
func (m *MockRepository) AddCard(ctx context.Context, deckId uuid.UUID, card *model.Card) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", ctx, deckId, card)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockRepositoryMockRecorder) AddCard(ctx, deckId, card any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockRepository)(nil).AddCard), ctx, deckId, card)
}

// AddDeck mocks base method.
func (m *MockRepository) AddDeck(ctx context.Context, deck *model.Deck) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeck", ctx, deck)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDeck indicates an expected call of AddDeck.
func (mr *MockRepositoryMockRecorder) AddDeck(ctx, deck any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeck", reflect.TypeOf((*MockRepository)(nil).AddDeck), ctx, deck)
}

// DeleteCard mocks base method.
func (m *MockRepository) DeleteCard(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockRepositoryMockRecorder) DeleteCard(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockRepository)(nil).DeleteCard), ctx, id)
}

// DeleteDeck mocks base method.
func (m *MockRepository) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeck", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeck indicates an expected call of DeleteDeck.
func (mr *MockRepositoryMockRecorder) DeleteDeck(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeck", reflect.TypeOf((*MockRepository)(nil).DeleteDeck), ctx, id)
}

// GetDeckById mocks base method.
func (m *MockRepository) GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeckById", ctx, id)
	ret0, _ := ret[0].(model.Deck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeckById indicates an expected call of GetDeckById.
func (mr *MockRepositoryMockRecorder) GetDeckById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeckById", reflect.TypeOf((*MockRepository)(nil).GetDeckById), ctx, id)
}

// UpdateCard mocks base method.
func (m *MockRepository) UpdateCard(ctx context.Context, id uuid.UUID, card dto.UpdateCardData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", ctx, id, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockRepositoryMockRecorder) UpdateCard(ctx, id, card any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockRepository)(nil).UpdateCard), ctx, id, card)
}

// UpdateDeck mocks base method.
func (m *MockRepository) UpdateDeck(ctx context.Context, id uuid.UUID, fieldsUpdate dto.UpdateDeckData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeck", ctx, id, fieldsUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeck indicates an expected call of UpdateDeck.
func (mr *MockRepositoryMockRecorder) UpdateDeck(ctx, id, fieldsUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeck", reflect.TypeOf((*MockRepository)(nil).UpdateDeck), ctx, id, fieldsUpdate)
}

// MockDeckRepository is a mock of DeckRepository interface.
type MockDeckRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeckRepositoryMockRecorder
	isgomock struct{}
}

// MockDeckRepositoryMockRecorder is the mock recorder for MockDeckRepository.
type MockDeckRepositoryMockRecorder struct {
	mock *MockDeckRepository
}

// NewMockDeckRepository creates a new mock instance.
func NewMockDeckRepository(ctrl *gomock.Controller) *MockDeckRepository {
	mock := &MockDeckRepository{ctrl: ctrl}
	mock.recorder = &MockDeckRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeckRepository) EXPECT() *MockDeckRepositoryMockRecorder {
	return m.recorder
}

// AddDeck mocks base method.
func (m *MockDeckRepository) AddDeck(ctx context.Context, deck *model.Deck) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeck", ctx, deck)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDeck indicates an expected call of AddDeck.
func (mr *MockDeckRepositoryMockRecorder) AddDeck(ctx, deck any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeck", reflect.TypeOf((*MockDeckRepository)(nil).AddDeck), ctx, deck)
}

// DeleteDeck mocks base method.
func (m *MockDeckRepository) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeck", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeck indicates an expected call of DeleteDeck.
func (mr *MockDeckRepositoryMockRecorder) DeleteDeck(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeck", reflect.TypeOf((*MockDeckRepository)(nil).DeleteDeck), ctx, id)
}

// GetDeckById mocks base method.
func (m *MockDeckRepository) GetDeckById(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeckById", ctx, id)
	ret0, _ := ret[0].(model.Deck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeckById indicates an expected call of GetDeckById.
func (mr *MockDeckRepositoryMockRecorder) GetDeckById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeckById", reflect.TypeOf((*MockDeckRepository)(nil).GetDeckById), ctx, id)
}

// UpdateDeck mocks base method.
func (m *MockDeckRepository) UpdateDeck(ctx context.Context, id uuid.UUID, fieldsUpdate dto.UpdateDeckData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeck", ctx, id, fieldsUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeck indicates an expected call of UpdateDeck.
func (mr *MockDeckRepositoryMockRecorder) UpdateDeck(ctx, id, fieldsUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeck", reflect.TypeOf((*MockDeckRepository)(nil).UpdateDeck), ctx, id, fieldsUpdate)
}

// MockCardRepository is a mock of CardRepository interface.
type MockCardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCardRepositoryMockRecorder
	isgomock struct{}
}

// MockCardRepositoryMockRecorder is the mock recorder for MockCardRepository.
type MockCardRepositoryMockRecorder struct {
	mock *MockCardRepository
}

// NewMockCardRepository creates a new mock instance.
func NewMockCardRepository(ctrl *gomock.Controller) *MockCardRepository {
	mock := &MockCardRepository{ctrl: ctrl}
	mock.recorder = &MockCardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardRepository) EXPECT() *MockCardRepositoryMockRecorder {
	return m.recorder
}

// AddCard mocks base method.
func (m *MockCardRepository) AddCard(ctx context.Context, deckId uuid.UUID, card *model.Card) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", ctx, deckId, card)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockCardRepositoryMockRecorder) AddCard(ctx, deckId, card any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockCardRepository)(nil).AddCard), ctx, deckId, card)
}

// DeleteCard mocks base method.
func (m *MockCardRepository) DeleteCard(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockCardRepositoryMockRecorder) DeleteCard(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockCardRepository)(nil).DeleteCard), ctx, id)
}

// UpdateCard mocks base method.
func (m *MockCardRepository) UpdateCard(ctx context.Context, id uuid.UUID, card dto.UpdateCardData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", ctx, id, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockCardRepositoryMockRecorder) UpdateCard(ctx, id, card any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockCardRepository)(nil).UpdateCard), ctx, id, card)
}
