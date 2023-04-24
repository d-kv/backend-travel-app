package repository

import (
	"errors"
)

var (
	ErrUUIDDuplicate    = errors.New("entity with the same uuid already exists")
	ErrUUIDNotPopulated = errors.New("field uuid must be populated")
)
