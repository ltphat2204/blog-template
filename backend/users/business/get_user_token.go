package business

import (
	"context"
	"time"

	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/users/entity"
	"golang.org/x/crypto/bcrypt"
)

type getTokenStorage interface {
	LoginUser(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
	UpdateLastAccess(ctx context.Context, condition map[string]interface{}) error
}

type getTokenBusiness struct {
	store getTokenStorage
}

func GetTokenFromUserBusiness(s getTokenStorage) *getTokenBusiness {
	return &getTokenBusiness{store: s}
}

func (biz *getTokenBusiness) GetTokenFromUser(ctx context.Context, username string, password string) (string, error) {
	user, err := biz.store.LoginUser(ctx, map[string]interface{}{"username": username})

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", entity.ErrNotFoundUsername
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", entity.ErrPswdNotMatch
	}

	var userNewAccess entity.UserAccess

	now := time.Now()
	userNewAccess.LastAccessAt = &now

	updateErr := biz.store.UpdateLastAccess(ctx, map[string]interface{}{"username": username})
	if updateErr != nil {
		return "", updateErr
	}

	token, tokenErr := common.GenerateToken(map[string]interface{}{"username": username})
 
	return token, tokenErr
}
