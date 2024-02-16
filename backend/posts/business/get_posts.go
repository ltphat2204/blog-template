package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

type getPostsStorage interface {
	GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error)
}

type getPostsBusiness struct {
	store getPostsStorage
}

func GetPostsBusiness(s getPostsStorage) *getPostsBusiness {
	return &getPostsBusiness{store: s}
}

func (biz *getPostsBusiness) GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error) {
	posts, err := biz.store.GetPosts(ctx, condition, pagination)

	return posts, err
}
