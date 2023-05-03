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
			Err(err).
			Str("method", mName).
			Msg("error from oAuthProvider")

		return "", err
	}

	err = c.tokenCache.SetUserID(ctx, rToken, uID)
	if err != nil {
		log.Error().
			Err(err).
			Str("method", mName).
			Msg("error from tokenCache")

		return "", err
	}

	return uID, nil
}
