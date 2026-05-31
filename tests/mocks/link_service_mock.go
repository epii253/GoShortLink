package mocks

import (
	contracts "project/internal/application/contracts"

	"github.com/stretchr/testify/mock"
)

type MockLinkService struct {
	mock.Mock
}

func (m *MockLinkService) AddNewLink(data contracts.LinkData) (*contracts.LinkAddResult, int) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Int(1)
	}
	return args.Get(0).(*contracts.LinkAddResult), args.Int(1)
}

func (m *MockLinkService) ExtractFullLink(data contracts.ShortLinkData) (*contracts.LinkExtractResult, int) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Int(1)
	}
	return args.Get(0).(*contracts.LinkExtractResult), args.Int(1)
}
