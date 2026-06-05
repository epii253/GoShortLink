package repositories

import (
	domain "project/internal/domain"
)

type ILinksRepo interface {
	CheckExsist(shrortLink string) (bool, error)

	TryAddItem(newItem *domain.Link) (bool, error)

	DeleteItemByShortLink(shrortLink string) (uint64, error)

	GetByLink(shortLink string) (*domain.Link, error)

	UpdateLink(Id uint, updatedLink *domain.Link) (bool, error)

	ClickCounterUp(Id uint) (*domain.Link, error)

	ClickCounterByShort(shortCode string) (*domain.Link, error)
}