package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

type createCategoryStorage interface {
	CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error)
	GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error)
}

type createCategoryBusiness struct {
	store createCategoryStorage
}

func CreateCategoryBusiness(s createCategoryStorage) *createCategoryBusiness {
	return &createCategoryBusiness{store: s}
}

func (biz *createCategoryBusiness) CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	oldCategory, err := biz.store.GetCategories(ctx, map[string]interface{}{"name": category.Name}, common.Pagination{}.GetDefault())
	if err != nil {
		return nil, err
	}

	if len(oldCategory) != 0 {
		return nil, entity.ErrNameExisting
	}

	createdCategory, createErr := biz.store.CreateCategory(ctx, category)

	return createdCategory, createErr
}
