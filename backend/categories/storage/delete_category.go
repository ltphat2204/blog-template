package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/categories/entity"
)

func (s *mySqlStorage) DeleteCategory(ctx context.Context, category *entity.Category) error {
	if err := s.storage.Delete(category).Error; err != nil {
		return err
	}

	return nil
}
