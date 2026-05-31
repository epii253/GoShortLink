package tests

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	contracts "project/internal/application/contracts"
	"project/internal/application/services/links"
	"project/internal/domain"
	"project/tests/mocks"
)

func TestAddNewLink_Success(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	repo.On("CheckExsist", mock.AnythingOfType("string")).Return(false, nil)
	repo.On("TryAddItem", mock.AnythingOfType("*domain.Link")).Return(true, nil)

	svc := links.NewLinkService(repo)
	result, status := svc.AddNewLink(contracts.LinkData{Url: "https://example.com"})

	assert.Equal(t, http.StatusCreated, status)
	assert.NotNil(t, result)
	assert.Len(t, result.ShortedUrl, 7)
	repo.AssertExpectations(t)
}

func TestAddNewLink_RepoError(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	repo.On("CheckExsist", mock.AnythingOfType("string")).Return(false, nil)
	repo.On("TryAddItem", mock.AnythingOfType("*domain.Link")).Return(false, errors.New("db error"))

	svc := links.NewLinkService(repo)
	result, status := svc.AddNewLink(contracts.LinkData{Url: "https://example.com"})

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Nil(t, result)
	repo.AssertExpectations(t)
}

func TestAddNewLink_Conflict(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	repo.On("CheckExsist", mock.AnythingOfType("string")).Return(false, nil)
	repo.On("TryAddItem", mock.AnythingOfType("*domain.Link")).Return(false, nil)

	svc := links.NewLinkService(repo)
	result, status := svc.AddNewLink(contracts.LinkData{Url: "https://example.com"})

	assert.Equal(t, http.StatusConflict, status)
	assert.Nil(t, result)
	repo.AssertExpectations(t)
}

func TestAddNewLink_SkipsExistingCode(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	repo.On("CheckExsist", mock.AnythingOfType("string")).Return(true, nil).Once()
	repo.On("CheckExsist", mock.AnythingOfType("string")).Return(false, nil).Once()
	repo.On("TryAddItem", mock.AnythingOfType("*domain.Link")).Return(true, nil)

	svc := links.NewLinkService(repo)
	result, status := svc.AddNewLink(contracts.LinkData{Url: "https://example.com"})

	assert.Equal(t, http.StatusCreated, status)
	assert.NotNil(t, result)
	repo.AssertExpectations(t)
}

func TestExtractFullLink_Success(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	expected := &domain.Link{ShortCode: "abc1234", FullUrl: "https://example.com"}
	repo.On("GetByLink", "abc1234").Return(expected, nil)

	svc := links.NewLinkService(repo)
	result, status := svc.ExtractFullLink(contracts.ShortLinkData{ShortLink: "abc1234"})

	assert.Equal(t, http.StatusFound, status)
	assert.NotNil(t, result)
	assert.Equal(t, "https://example.com", result.FullUrl)
	repo.AssertExpectations(t)
}

func TestExtractFullLink_RepoError(t *testing.T) {
	repo := new(mocks.MockLinksRepo)
	repo.On("GetByLink", "abc1234").Return(nil, errors.New("connection refused"))

	svc := links.NewLinkService(repo)
	result, status := svc.ExtractFullLink(contracts.ShortLinkData{ShortLink: "abc1234"})

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Nil(t, result)
	repo.AssertExpectations(t)
}
