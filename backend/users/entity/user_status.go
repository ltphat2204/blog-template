package entity

import (
	"errors"
)

var (
	ErrPswdCnntHash = errors.New("password can not be hashed")
	ErrExtUsername = errors.New("username exitsted")
	ErrNotFoundUsername = errors.New("username not found")
	ErrPswdNotMatch = errors.New("password does not match")
	ErrDelAdmin = errors.New("admin account can not be deleted")
)
