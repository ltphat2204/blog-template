package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/subcribers/entity"
)

type getSubcribersStorage interface {
	GetSubcribers(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Subcriber, error)
}

type getSubcribersBusiness struct {
	store getSubcribersStorage
}

func GetSubcribersBusiness(s getSubcribersStorage) *getSubcribersBusiness {
	return &getSubcribersBusiness{store: s}
}

func (biz *getSubcribersBusiness) GetSubcribers(ctx context.Context, condition map[string]interface{}, pagination *common.Pagination) ([]entity.Subcriber, error) {
	subcribers, err := biz.store.GetSubcribers(ctx, condition, pagination)

	return subcribers, err
}
