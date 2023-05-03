package ictrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"

	"github.com/d-kv/backend-travel-app/pkg/user-service/cache"
	ctrl_v0 "github.com/d-kv/backend-travel-app/pkg/user-service/controller/v0"
	"github.com/rs/zerolog/log"
)

func (c *UserController) Authenticate(ctx context.Context, refreshToken string) (string, error) {
	const mName = "UserController.Authenticate"

	uID, err := c.tokenCache.UserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, cache.ErrRefreshTokenNotFound) {
			log.Warn().
				Str("method", mName).
				Err(err).
				Msg("error from tokenCache")

			return "", ctrl_v0.ErrRefreshTokenNotFound
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from tokenCache")

		return "", err
	}
	return uID, nil
}
