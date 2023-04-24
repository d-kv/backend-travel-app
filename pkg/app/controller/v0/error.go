package controllerv0

import (
	"errors"
)

var (
	ErrCategoryNotSpecified     = errors.New("category not specified")
	ErrNoPlaces                 = errors.New("no places with given criteria")
	ErrVersionNotCompatible     = errors.New("version not compatible")
	ErrUserIsBlocked            = errors.New("user is blocked")
	ErrUserNotFound             = errors.New("user not found")
	ErrRefreshTokenNotFound     = errors.New("refresh token not found")
	ErrAchievementAlreadyExists = errors.New("achievement already exist")
	ErrBadAchievement           = errors.New("bad achievement")
)
