package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/entity"
)

type getPostStorage interface {
	GetPosts(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.PostDisplay, error)
}

type getPostBusiness struct {
	store getPostStorage
}

func GetPostBusiness(s getPostStorage) *getPostBusiness {
	return &getPostBusiness{store: s}
}

func (biz *getPostBusiness) GetPost(ctx context.Context, id uint) (*entity.PostDisplay, error) {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetPosts(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return nil, err
	}

	if len(post) == 0 {
		return nil, entity.ErrPostNotFound
	}

	return &post[0], nil
}
