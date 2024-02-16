package entity

import "errors"

var (
	ErrTitleExisting = errors.New("Title is existing!")
	ErrPostNotFound = errors.New("Post not found!")
	ErrPostDeleted = errors.New("Post is deleted!")
)
