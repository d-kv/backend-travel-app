package repository

import (
	"errors"
)

var (
	ErrUUIDNotPopulated = errors.New("field uuid must be populated")
)
