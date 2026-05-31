package repositories

import (
	domain "project/internal/domain"
)

type ILinksRepo interface {
	CheckExsist(shrortLink string) (bool, error)

	TryAddItem(newItem *domain.Link) (bool, error)

	DeleteItemByShortLink(shrortLink string) (bool, error)

	GetByLink(shortLink string) (*domain.Link, error)
}