package links

import (
	"net/http"
	"project/internal/application/repositories"
	"project/internal/application/services/utilities"
)

type ILinkService interface {
	AddNewLink(data LinkData) (*LinkAddResult, int)
	ExtractFullLink(data ShortLinkData) (*LinkExtractResult, int)
}

type LinkService struct {
	repo   repositories.ILinksRepo
	urlLen int
}

func NewLinkService(repo repositories.ILinksRepo) *LinkService {
	return &LinkService{repo: repo, urlLen: 7}
}

func (service *LinkService) AddNewLink(data LinkData) (*LinkAddResult, int) {
	var candidate string = utilities.RandomCode(service.urlLen)

	for ; service.repo.CheckExsist(candidate); {
		candidate = utilities.RandomCode(service.urlLen)
	}
	service.repo.TryAddItem(data.Url, candidate)

	result := &LinkAddResult{ShortedUrl: candidate}

	return result, http.StatusCreated
}

func (service *LinkService) ExtractFullLink(data ShortLinkData) (*LinkExtractResult, int) {
	fullUrl, ok := service.repo.GetByLink(data.ShortLink)

	if !ok {
		return nil, http.StatusNotFound
	}
	return &LinkExtractResult{FullUrl: *fullUrl}, http.StatusFound
}
