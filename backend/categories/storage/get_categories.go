package storage

import (
	"context"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

func (s *mySqlStorage) GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error) {
	var categories []entity.Category

	if err := s.storage.Table("categories").Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	if err := s.storage.Offset((pagination.Page-1)*pagination.Limit).Limit(pagination.Limit).Find(&categories, condition).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
