package gateway

import "errors"

var (
	ErrTokenIsExpired = errors.New("access token is expired")
)
