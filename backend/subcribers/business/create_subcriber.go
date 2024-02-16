package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/subcribers/entity"
)

type createSubcriberStorage interface {
	CreateSubcriber(ctx context.Context, post *entity.Subcriber) (*entity.Subcriber, error)
	GetSubcribers(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Subcriber, error)
}

type createSubcriberBusiness struct {
	store createSubcriberStorage
}

func CreateSubcriberBusiness(s createSubcriberStorage) *createSubcriberBusiness {
	return &createSubcriberBusiness{store: s}
}

func (biz *createSubcriberBusiness) CreateSubcriber(ctx context.Context, post *entity.Subcriber) (*entity.Subcriber, error) {
	oldSubcriber, err := biz.store.GetSubcribers(ctx, map[string]interface{}{"email": post.Email}, common.Pagination{}.GetDefault())
	if err != nil {
		return nil, err
	}

	if len(oldSubcriber) != 0 {
		return nil, entity.ErrEmailExisting
	}

	createdSubcriber, createErr := biz.store.CreateSubcriber(ctx, post)

	return createdSubcriber, createErr
}
