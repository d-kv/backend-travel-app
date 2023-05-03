package place_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import "errors"

var (
	ErrCategoryNotSpecified = errors.New("category not specified")
	ErrNoPlaces             = errors.New("no places with given criteria")
)
