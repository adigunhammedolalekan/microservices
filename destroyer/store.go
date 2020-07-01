package destroyer

import (
	"github.com/adigunhammedolalekan/microservices-test/db"
	"github.com/adigunhammedolalekan/microservices-test/types"
	"github.com/jinzhu/gorm"
	"os"
)

type Store interface {
	ListTargets() ([]*types.Target, error)
}

type databaseStore struct {
	db *gorm.DB
}

func NewStore() (Store, error) {
	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &databaseStore{db: database}, nil
}

func (s *databaseStore) ListTargets() ([]*types.Target, error) {
	data := make([]*types.Target, 0)
	err := s.db.Table("events").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
