package business

import (
	"context"
	"time"

	"github.com/ltphat2204/blog/backend/users/entity"
	"golang.org/x/crypto/bcrypt"
)

type createUserStorage interface {
	CreateUser(ctx context.Context, data *entity.User) error
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

type createUserBusiness struct {
	store createUserStorage
}

func CreateNewUserBusiness(s createUserStorage) *createUserBusiness {
	return &createUserBusiness{store: s}
}

func (biz *createUserBusiness) CreateNewUser(ctx context.Context, data *entity.User) error {
	if oldUser, _ := biz.store.GetUser(ctx, map[string]interface{}{"username": data.Username}); oldUser != nil {
		return entity.ErrExtUsername
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.ErrPswdCnntHash
	}

	data.Password = string(hashedPassword)

	now := time.Now()
	data.CreatedAt = &now
	data.LastAccessAt = &now

	err = biz.store.CreateUser(ctx, data)

	return err
}
