package storage

import (
	"context"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

func (s *mySqlStorage) CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	if err := s.storage.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
