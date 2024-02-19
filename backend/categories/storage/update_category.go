package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/categories/entity"
)

func (s *mySqlStorage) UpdateCategory(ctx context.Context, condition map[string]interface{}, category *entity.Category) error {
	if err := s.storage.Where(condition).Updates(category).Error; err != nil {
		return err
	}

	return nil
}
