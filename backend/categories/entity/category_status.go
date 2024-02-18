package entity

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrNameExisting = errors.New("this category has created already")
)