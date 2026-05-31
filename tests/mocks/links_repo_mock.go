package mocks

import (
	"project/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockLinksRepo struct {
	mock.Mock
}

func (m *MockLinksRepo) CheckExsist(shortLink string) (bool, error) {
	args := m.Called(shortLink)
	return args.Bool(0), args.Error(1)
}

func (m *MockLinksRepo) TryAddItem(newItem *domain.Link) (bool, error) {
	args := m.Called(newItem)
	return args.Bool(0), args.Error(1)
}

func (m *MockLinksRepo) DeleteItemByShortLink(shortLink string) (bool, error) {
	args := m.Called(shortLink)
	return args.Bool(0), args.Error(1)
}

func (m *MockLinksRepo) GetByLink(shortLink string) (*domain.Link, error) {
	args := m.Called(shortLink)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Link), args.Error(1)
}
