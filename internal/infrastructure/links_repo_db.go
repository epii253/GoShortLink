package infrastructure

import (
	"project/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type LinksDbRepo struct {
	db *gorm.DB
}

func NewLinksDbRepo(db *gorm.DB) *LinksDbRepo {
	return &LinksDbRepo{db: db}
}

func (repo *LinksDbRepo) CheckExsist(shrortLink string) (bool, error) {
	var count int64 = 0
	result := repo.db.
		Model(&domain.Link{}).
		Where("short_code = ?", shrortLink).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (repo *LinksDbRepo) TryAddItem(newItem *domain.Link) (bool, error) {
	result := repo.db.Create(&newItem)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (repo *LinksDbRepo) DeleteItemByShortLink(shrortLink string) (bool, error) {
	result := repo.db.
		Model(&domain.Link{}).
		Where("short_link = ?", shrortLink).
		Delete(&domain.Link{})

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (repo *LinksDbRepo) GetByLink(shortLink string) (*domain.Link, error) {
	var link domain.Link

	result := repo.db.
		Model(&domain.Link{}).
		Where("short_code = ?", shortLink).
		First(&link)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}
	
	return &link, nil
}