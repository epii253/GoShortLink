package domain

import (
	gorm "gorm.io/gorm"
)

type Link struct {
	gorm.Model

	ShortCode string `gorm:"uniqueIndex;not null"`
	FullUrl string `gorm:"not null"`
}

func NewLink(shortLink string, fullLink string) (*Link, error) {
	return &Link{
		ShortCode: shortLink,
		FullUrl: fullLink,
	}, nil
}