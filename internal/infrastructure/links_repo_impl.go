package infrastructure

import "sync"

type LinksInMemoryRepo struct {
	mux sync.Mutex

	data map[string]string // short: full
}

func NewLinksInMemoryRepo() *LinksInMemoryRepo {
	return &LinksInMemoryRepo{mux: sync.Mutex{}, data: map[string]string{}}
}

func (repo *LinksInMemoryRepo) CheckExsist(shrortLink string) bool {
	repo.mux.Lock()

	defer repo.mux.Unlock()

	if _, ok := repo.data[shrortLink]; !ok {
		return false
	} else {
		return true
	}
}

func (repo *LinksInMemoryRepo) TryAddItem(fullLink string, shrortLink string) bool {
	repo.mux.Lock()

	defer repo.mux.Unlock()

	if repo.CheckExsist(shrortLink) {
		return false
	}

	repo.data[shrortLink] = fullLink
	return true
}

func (repo *LinksInMemoryRepo) DeleteItem(shrortLink string) {
	repo.mux.Lock()

	delete(repo.data, shrortLink)

	repo.mux.Unlock()
}

func (repo *LinksInMemoryRepo) GetByLink(shrortLink string) (*string, bool) {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	link, ok := repo.data[shrortLink]
	return &link, ok
}
