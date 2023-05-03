package repository

import "errors"

var (
	ErrPlaceNotFound    = errors.New("place not found")
	ErrUUIDNotPopulated = errors.New("field uuid must be populated")
)
