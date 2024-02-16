package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

func (s *mySqlStorage) GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error) {
	var posts []entity.PostDisplay

	if err := s.storage.Table("posts").Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	if err := s.storage.Offset((pagination.Page-1)*pagination.Limit).Limit(pagination.Limit).Find(&posts, condition).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
