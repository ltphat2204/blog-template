package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/posts/entity"
)

func (s *mySqlStorage) UpdatePost(ctx context.Context, condition map[string]interface{}, post *entity.PostUpdate) error {
	if err := s.storage.Where(condition).Updates(post).Error; err != nil {
		return err
	}

	return nil
}
