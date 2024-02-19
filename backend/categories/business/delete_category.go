package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

type deleteCategoryStorage interface {
	GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error)
	DeleteCategory(ctx context.Context, category *entity.Category) error
}

type deleteCategoryBusiness struct {
	store deleteCategoryStorage
}

func DeleteCategoryBusiness(s deleteCategoryStorage) *deleteCategoryBusiness {
	return &deleteCategoryBusiness{store: s}
}

func (biz *deleteCategoryBusiness) DeleteCategory(ctx context.Context, id uint) error {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetCategories(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return err
	}

	if len(post) == 0 {
		return entity.ErrCategoryNotFound
	}

	err = biz.store.DeleteCategory(ctx, &post[0])

	return err
}
