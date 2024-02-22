package business

import (
	"context"

	"github.com/ltphat2204/blog/backend/users/entity"
)

type findUserStorage interface {
	GetAllUsers(ctx context.Context, condition map[string]interface{}) ([]entity.UserShow, error)
}

type findUserBusiness struct {
	store findUserStorage
}

func FindAllUsersBusiness(s findUserStorage) *findUserBusiness {
	return &findUserBusiness{store: s}
}

func (biz *findUserBusiness) FindAllUsers(ctx context.Context, condition map[string]interface{}) ([]entity.UserShow, error) {
	user, err := biz.store.GetAllUsers(ctx, condition)

	return user, err
}
