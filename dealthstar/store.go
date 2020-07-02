package dealthstar

import (
	"github.com/adigunhammedolalekan/microservices-test/db"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/jinzhu/gorm"
	"os"
)

type Store interface {
	CreateEvents(*types.Event) error
	Close() error
}

type databaseStore struct {
	db *gorm.DB
}

func (s *databaseStore) Close() error {
	return s.db.Close()
}

func (s *databaseStore) CreateEvents(event *types.Event) error {
	for _, target := range event.Data {
		if err := s.db.Create(target).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewStore() (Store, error) {
	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &databaseStore{db: database}, nil
}
