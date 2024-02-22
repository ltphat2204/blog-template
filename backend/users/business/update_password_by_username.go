package business

import (
	"context"
	"time"

	"github.com/ltphat2204/blog/backend/users/entity"
	"golang.org/x/crypto/bcrypt"
)

type updateUserStorage interface {
	UpdateUser(ctx context.Context, data *entity.User) error
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

type updateUserBusiness struct {
	store updateUserStorage
}

func UpdatePasswordUserBusiness(s updateUserStorage) *updateUserBusiness {
	return &updateUserBusiness{store: s}
}

func (biz *updateUserBusiness) UpdatePasswordUser(ctx context.Context, username string, password string) error {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"username": username})

	if err != nil {
		return err
	}

	if user == nil {
		return entity.ErrNotFoundUsername
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.ErrPswdCnntHash
	}

	var userWithNewPassword entity.User
	userWithNewPassword = entity.CopyInformationFrom(*user)
	userWithNewPassword.Password = string(hashedPassword)

	now := time.Now()
	userWithNewPassword.LastAccessAt = &now

	err = biz.store.UpdateUser(ctx, &userWithNewPassword)

	return err
}
