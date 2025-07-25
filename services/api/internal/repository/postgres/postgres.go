package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
