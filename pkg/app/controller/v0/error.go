package controller_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"errors"
)

var (
	ErrVersionNotCompatible = errors.New("version not compatible")
)
