package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import "errors"

var (
	errLatLngCoupling = errors.New("lat & lng are required")
	errInvalidLatLng  = errors.New("invalid ll parameter")

	errMinDMaxDCoupling    = errors.New("min_d & min_d must either be both present or both absent")
	errInvalidMinD         = errors.New("invalid min_d parameter")
	errInvalidMaxD         = errors.New("invalid max_d parameter")
	errMaxDSmallerThanMinD = errors.New("max_d must be greater than or equal to min_d")

	errInvalidSkipN = errors.New("invalid skip_n parameter")
	errInvalidResN  = errors.New("invalid res_n parameter")

	errInvalidBody = errors.New("invalid categories")
)
