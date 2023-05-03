package controllerv0

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (c *Controller) AuthorizeOAuthUser(ctx context.Context, aToken, rToken string) (string, error) {
	const mName = "Controller.AuthorizeOAuthUser"

	uID, err := c.oAuthProvider.UserID(ctx, aToken)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from oAuthProvider")

		return "", err
	}

	err = c.tokenCache.SetUserID(ctx, rToken, uID)
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from tokenCache")

		return "", err
	}

	return uID, nil
}
