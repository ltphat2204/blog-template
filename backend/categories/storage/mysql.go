package storage

import (
	"gorm.io/gorm"
)

type mySqlStorage struct {
	storage *gorm.DB
}

func NewMySqlStorage(db *gorm.DB) *mySqlStorage {
	return &mySqlStorage{storage: db}
}
