package db

import (
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/jinzhu/gorm"
)

func Connect(url string) (*gorm.DB, error) {
	database, err := gorm.Open(url)
	if err != nil {
		return nil, err
	}
	if err := database.DB().Ping(); err != nil {
		return nil, err
	}
	runMigration(database)
	return database, nil
}

func runMigration(database *gorm.DB) {
	database.AutoMigrate(&types.Event{})
}
