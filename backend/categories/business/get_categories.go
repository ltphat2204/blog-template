package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/entity"
)

type getCategoriesStorage interface {
	GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error)
}

type getCategoriesBusiness struct {
	store getCategoriesStorage
}

func GetCategoriesBusiness(s getCategoriesStorage) *getCategoriesBusiness {
	return &getCategoriesBusiness{store: s}
}

func (biz *getCategoriesBusiness) GetCategories(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Category, error) {
	categories, err := biz.store.GetCategories(ctx, condition, pagination)

	return categories, err
}
