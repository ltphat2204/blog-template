package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

type updatePostStorage interface {
	GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error)
	UpdatePost(ctx context.Context, condition map[string]interface{}, post *entity.PostUpdate) error
}

type updatePostBusiness struct {
	store updatePostStorage
}

func UpdatePostBusiness(s updatePostStorage) *updatePostBusiness {
	return &updatePostBusiness{store: s}
}

func (biz *updatePostBusiness) UpdatePost(ctx context.Context, id uint, newPost *entity.PostUpdate) error {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetPosts(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return err
	}

	if len(post) == 0 {
		return entity.ErrPostNotFound
	}

	err = biz.store.UpdatePost(ctx, idFilter, newPost)

	return err
}
