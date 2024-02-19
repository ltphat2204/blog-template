package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/categories/entity"
	"github.com/ltphat2204/blog/backend/common"
)

type updateCategoryStorage interface {
	GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, condition map[string]interface{}, category *entity.Category) error
}

type updateCategoryBusiness struct {
	store updateCategoryStorage
}

func UpdateCategoryBusiness(s updateCategoryStorage) *updateCategoryBusiness {
	return &updateCategoryBusiness{store: s}
}

func (biz *updateCategoryBusiness) UpdateCategory(ctx context.Context, id uint, newCategory *entity.Category) error {
	idFilter := map[string]interface{}{"id": id}
	post, err := biz.store.GetCategories(ctx, idFilter, common.Pagination{}.GetDefault())

	if err != nil {
		return err
	}

	if len(post) == 0 {
		return entity.ErrCategoryNotFound
	}

	err = biz.store.UpdateCategory(ctx, idFilter, newCategory)

	return err
}
