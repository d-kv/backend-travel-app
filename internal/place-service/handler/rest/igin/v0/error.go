package igin_v0

import "errors"

var (
	ErrLatLngNotFound = errors.New("missing ll parameter")
	ErrLatLngNotValid = errors.New("invalid ll parameter")
)
