package infrastructure

import (
	"errors"
	"sync"

	domain "project/internal/domain"
)

type LinksInMemoryRepo struct {
	mux  sync.Mutex
	data map[string]*domain.Link // shortCode: Link
}

func NewLinksInMemoryRepo() *LinksInMemoryRepo {
	return &LinksInMemoryRepo{data: map[string]*domain.Link{}}
}

func (repo *LinksInMemoryRepo) CheckExsist(shrortLink string) (bool, error) {
	_, ok := repo.data[shrortLink]
	return ok, nil
}

func (repo *LinksInMemoryRepo) TryAddItem(newItem *domain.Link) (bool, error) {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	if _, ok := repo.data[newItem.ShortCode]; ok {
		return false, nil
	}

	repo.data[newItem.ShortCode] = newItem
	return true, nil
}

func (repo *LinksInMemoryRepo) DeleteItemByShortLink(shrortLink string) (bool, error) {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	if _, ok := repo.data[shrortLink]; !ok {
		return false, nil
	}

	delete(repo.data, shrortLink)
	return true, nil
}

func (repo *LinksInMemoryRepo) GetByLink(shortLink string) (*domain.Link, error) {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	link, ok := repo.data[shortLink]
	if !ok {
		return nil, errors.New("link not found")
	}
	return link, nil
}
