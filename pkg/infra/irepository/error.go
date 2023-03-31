package irepository

import "errors"

// PlaceStore errors.
var (
	ErrPlaceNotFound = errors.New("place is not found")
)

// PlaceStore errors.
var (
	ErrUserNotFound = errors.New("user is not found")
)

// Common errors.
var (
	ErrUUIDDuplicate    = errors.New("entity with the same uuid already exists")
	ErrUUIDNotPopulated = errors.New("field uuid must be populated")
)
