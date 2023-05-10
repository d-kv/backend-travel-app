package ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"errors"
)

var (
	ErrUserIsBlocked            = errors.New("user is blocked")
	ErrUserNotFound             = errors.New("user not found")
	ErrRefreshTokenNotFound     = errors.New("refresh token not found")
	ErrAchievementAlreadyExists = errors.New("achievement already exist")
	ErrBadAchievement           = errors.New("bad achievement")
	ErrVersionNotCompatible     = errors.New("version not compatible")
)
