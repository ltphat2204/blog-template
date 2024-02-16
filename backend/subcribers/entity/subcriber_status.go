package entity

import "errors"

var (
	ErrEmailExisting = errors.New("This email has subcribed already!")
)
