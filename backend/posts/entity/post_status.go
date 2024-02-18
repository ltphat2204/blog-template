package entity

import "errors"

var (
	ErrTitleExisting = errors.New("title is existing")
	ErrPostNotFound = errors.New("post not found")
	ErrPostDeleted = errors.New("post is deleted")
)
