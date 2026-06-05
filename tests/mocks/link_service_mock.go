package mocks

import (
	contracts "project/internal/application/contracts"

	"github.com/stretchr/testify/mock"
)

type MockLinkService struct {
	mock.Mock
}

func (m *MockLinkService) AddNewLink(data contracts.LinkData) (*contracts.LinkAddResponse, int) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Int(1)
	}
	return args.Get(0).(*contracts.LinkAddResponse), args.Int(1)
}

func (m *MockLinkService) ExtractFullLink(data contracts.ShortLinkRequest) (*contracts.LinkExtractResponse, int) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Int(1)
	}
	return args.Get(0).(*contracts.LinkExtractResponse), args.Int(1)
}
