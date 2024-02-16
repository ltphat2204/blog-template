package storage

import (
	"context"
	"github.com/ltphat2204/blog/backend/subcribers/entity"
)

func (s *mySqlStorage) CreateSubcriber(ctx context.Context, subcriber *entity.Subcriber) (*entity.Subcriber, error) {
	if err := s.storage.Create(subcriber).Error; err != nil {
		return nil, err
	}

	return subcriber, nil
}