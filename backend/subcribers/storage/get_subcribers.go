package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/subcribers/entity"
)

func (s *mySqlStorage) GetSubcribers(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Subcriber, error) {
	var subcribers []entity.Subcriber

	if err := s.storage.Table("subcribers").Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	if err := s.storage.Offset((pagination.Page-1)*pagination.Limit).Limit(pagination.Limit).Find(&subcribers, condition).Error; err != nil {
		return nil, err
	}

	return subcribers, nil
}
