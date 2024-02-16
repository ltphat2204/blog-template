package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

type createPostStorage interface {
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error)
}

type createPostBusiness struct {
	store createPostStorage
}

func CreatePostBusiness(s createPostStorage) *createPostBusiness {
	return &createPostBusiness{store: s}
}

func (biz *createPostBusiness) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	oldPost, err := biz.store.GetPosts(ctx, map[string]interface{}{"title": post.Title}, common.Pagination{}.GetDefault())
	if err != nil {
		return nil, err
	}

	if len(oldPost) != 0 {
		return nil, entity.ErrTitleExisting
	}

	createdPost, createErr := biz.store.CreatePost(ctx, post)

	return createdPost, createErr
}
