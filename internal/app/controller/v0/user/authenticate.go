package iuser_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	user_controller_v0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0/user"
	tokencache "github.com/d-kv/backend-travel-app/pkg/infra/cache/token"
)

func (c *UserController) Authenticate(ctx context.Context, refreshToken string) (string, error) {
	const mName = "UserController.Authenticate"

	uID, err := c.tokenCache.UserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, tokencache.ErrRefreshTokenNotFound) {
			log.Warn().
				Str("method", mName).
				Err(err).
				Msg("error from tokenCache")

			return "", user_controller_v0.ErrRefreshTokenNotFound
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from tokenCache")

		return "", err
	}
	return uID, nil
}
