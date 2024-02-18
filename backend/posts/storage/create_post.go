package storage

import (
	"context"
	"errors"

	"github.com/ltphat2204/blog/backend/posts/entity"
)

func (s *mySqlStorage) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	var checkCate int64
	if s.storage.Table("categories").Where("id = ?", post.CategoryID).Count(&checkCate); checkCate == 0 {
		return nil, errors.New("category not found")
	}

	if err := s.storage.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
