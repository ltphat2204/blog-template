package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

type getCategoryStorage interface {
	GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error)
}

type getCategoryBusiness struct {
	store getCategoryStorage
}

func GetCategoryBusiness(s getCategoryStorage) *getCategoryBusiness {
	return &getCategoryBusiness{store: s}
}

func (biz *getCategoryBusiness) GetCategory(ctx context.Context, id uint) (*entity.Category, error) {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetCategories(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return nil, err
	}

	if len(post) == 0 {
		return nil, entity.ErrCategoryNotFound
	}

	return &post[0], nil
}
