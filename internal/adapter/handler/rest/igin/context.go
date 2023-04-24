package ggonic

import (
	"github.com/gin-gonic/gin"
)

const (
	userUUID = 0
)

var contextKeys = map[int32]string{ //nolint:gochecknoglobals // Not variable but constant
	userUUID: "user_uuid",
}

func SetUserUUID(c *gin.Context, id string) error {
	_, err := UserUUID(c)
	if err != nil {
		return ErrAlreadyInContext
	}

	c.Set(contextKeys[userUUID], id)
	return nil
}

func UserUUID(c *gin.Context) (string, error) {
	if c.Value(userUUID) == nil {
		return "", ErrNotInContext
	}

	return c.Value(userUUID).(string), nil
}
