package icontrollerv0

import "errors"

var (
	ErrCategoryNotSpecified = errors.New("category not specified")
	ErrNoPlaces             = errors.New("no places with given category")
	ErrVersionNotCompatible = errors.New("version not compatible")
)
