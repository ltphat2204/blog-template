package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/users/entity"
)

type getUserStorage interface {
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

type getUserBusiness struct {
	store getUserStorage
}

func GetUserByUsernameBusiness(s getUserStorage) *getUserBusiness {
	return &getUserBusiness{store: s}
}

func (biz *getUserBusiness) GetUserByUsername(ctx context.Context, username string) (*entity.UserShow, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"username": username})

	return user, err
}
