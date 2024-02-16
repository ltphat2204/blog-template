package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

type deletePostStorage interface {
	GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error)
	DeletePost(ctx context.Context, post *entity.PostDisplay) error
}

type deletePostBusiness struct {
	store deletePostStorage
}

func DeletePostBusiness(s deletePostStorage) *deletePostBusiness {
	return &deletePostBusiness{store: s}
}

func (biz *deletePostBusiness) DeletePost(ctx context.Context, id uint) error {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetPosts(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return err
	}

	if len(post) == 0 {
		return entity.ErrPostNotFound
	}

	err = biz.store.DeletePost(ctx, &post[0])

	return err
}
