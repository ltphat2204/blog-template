package storage

import (
	"context"

	"github.com/ltphat2204/blog/backend/posts/entity"
)

func (s *mySqlStorage) DeletePost(ctx context.Context, post *entity.PostDisplay) error {
	if err := s.storage.Delete(post).Error; err != nil {
		return err
	}

	return nil
}
