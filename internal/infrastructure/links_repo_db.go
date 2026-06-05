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

func (repo *LinksDbRepo) DeleteItemByShortLink(shrortLink string) (uint64, error) {
	var link domain.Link

	result := repo.db.
		Model(&domain.Link{}).
		Where("short_code = ?", shrortLink).
		First(&link)

	if result.Error != nil {
		return 0, result.Error
	}

	if errDelete := repo.db.Delete(&link).Error; errDelete != nil {
    	return 0, errDelete
	}

	return link.Clicks, nil
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
	
	updatedLink, errUpdate := repo.ClickCounterUp(link.ID)

	if errUpdate != nil {
		return nil, errUpdate
	}
	
	return updatedLink, nil
}

func (repo *LinksDbRepo) UpdateLink(Id uint, updatedLink *domain.Link) (bool, error) {
	result := repo.db.
		Where("id = ?", Id).
		Updates(updatedLink)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (repo *LinksDbRepo) ClickCounterUp(Id uint) (*domain.Link, error) {
	var updatedLink domain.Link
	
	result := repo.db.
		Model(&updatedLink).
		Where("id = ?", Id).
		Update("clicks", gorm.Expr("clicks + ?", 1))

	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedLink, nil
}

func (repo *LinksDbRepo) ClickCounterByShort(shortCode string) (*domain.Link, error) {
	var updatedLink domain.Link
	
	result := repo.db.
		Where("short_code = ?", shortCode).
		First(&updatedLink)

	if result.Error != nil {
		return nil, result.Error
	}

	return repo.ClickCounterUp(updatedLink.ID)
}