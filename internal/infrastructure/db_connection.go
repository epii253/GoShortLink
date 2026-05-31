package infrastructure

import (
	"fmt"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	domain "project/internal/domain"
	settings "project/internal/settings"
)

func NewPsqlDB(cnf *settings.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cnf.DBHost,
		cnf.DBUser,
		cnf.DBPass,
		cnf.DBName,
		cnf.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&domain.Link{}); err != nil {
		return nil, err
	}

	return db, nil
}