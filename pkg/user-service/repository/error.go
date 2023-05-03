package repository

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUUIDDuplicate    = errors.New("entity with the same uuid already exists")
	ErrUUIDNotPopulated = errors.New("field uuid must be populated")
)
