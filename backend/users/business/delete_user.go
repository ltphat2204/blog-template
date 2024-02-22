package business

import (
	"context"
	"os"

	"github.com/ltphat2204/blog/backend/users/entity"
)

type deleteUserStorage interface {
	DeleteUser(ctx context.Context, username string) error
}

type deleteUserBusiness struct {
	store deleteUserStorage
}

func DeleteUserBusiness(s deleteUserStorage) *deleteUserBusiness {
	return &deleteUserBusiness{store: s}
}

func (biz *deleteUserBusiness) DeleteUser(ctx context.Context, username string) error {
	if username == os.Getenv("ADMIN_USERNAME") {
		return entity.ErrDelAdmin
	}

	err := biz.store.DeleteUser(ctx, username)

	return err
}
