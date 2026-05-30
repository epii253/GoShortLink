package links

import (
	"net/http"
	"project/internal/application/repositories"
	"project/internal/application/services/utilities"
	contracts "project/internal/application/contracts"
)

type ILinkService interface {
	AddNewLink(data contracts.LinkData) (*contracts.LinkAddResult, int)
	ExtractFullLink(data contracts.ShortLinkData) (*contracts.LinkExtractResult, int)
}

type LinkService struct {
	repo   repositories.ILinksRepo
	urlLen int
}

func NewLinkService(repo repositories.ILinksRepo) *LinkService {
	return &LinkService{repo: repo, urlLen: 7}
}

func (service *LinkService) AddNewLink(data contracts.LinkData) (*contracts.LinkAddResult, int) {
	var candidate string = utilities.RandomCode(service.urlLen)

	for service.repo.CheckExsist(candidate) {
		candidate = utilities.RandomCode(service.urlLen)
	}
	service.repo.TryAddItem(data.Url, candidate)
	
	result := &contracts.LinkAddResult{ShortedUrl: candidate}

	return result, http.StatusCreated
}

func (service *LinkService) ExtractFullLink(data contracts.ShortLinkData) (*contracts.LinkExtractResult, int) {
	fullUrl, ok := service.repo.GetByLink(data.ShortLink)

	if !ok {
		return nil, http.StatusNotFound
	}
	return &contracts.LinkExtractResult{FullUrl: *fullUrl}, http.StatusFound
}
