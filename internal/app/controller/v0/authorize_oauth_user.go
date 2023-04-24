package controllerv0

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (c *Controller) AuthorizeOAuthUser(ctx context.Context, aToken, rToken string) (string, error) {
	uID, err := c.oAuthProvider.UserID(ctx, aToken)
	if err != nil {
		log.Warn().
			Err(err)
		return "", err
	}

	err = c.TokenCache.SetUserID(ctx, rToken, uID)
	if err != nil {
		log.Error().
			Err(err)
		return "", err
	}

	return uID, nil
}
