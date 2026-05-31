package links

import (
	"net/http"
	"project/internal/application/repositories"
	"project/internal/application/services/utilities"

	domain "project/internal/domain"
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
	var candidateCode string = utilities.RandomCode(service.urlLen)

	for range 10 {
		exsist, _ := service.repo.CheckExsist(candidateCode)

		if exsist == false {
			break
		}

		candidateCode = utilities.RandomCode(service.urlLen)
	}

	candidateLink, _ := domain.NewLink(candidateCode, data.Url)

	isAdded, err := service.repo.TryAddItem(candidateLink)
	
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if !isAdded {
		return nil, http.StatusConflict
	}

	result := &contracts.LinkAddResult{ShortedUrl: candidateLink.ShortCode}

	return result, http.StatusCreated
}

func (service *LinkService) ExtractFullLink(data contracts.ShortLinkData) (*contracts.LinkExtractResult, int) {
	link, err := service.repo.GetByLink(data.ShortLink)

	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if link == nil {
		return nil, http.StatusNotFound
	}
	
	return &contracts.LinkExtractResult{FullUrl: link.FullUrl}, http.StatusFound
}
