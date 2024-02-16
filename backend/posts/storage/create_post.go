package storage

import (
	"context"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

func (s *mySqlStorage) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	if err := s.storage.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
