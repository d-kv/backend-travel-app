package controllerv0

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	icontrollerv0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0"
	tokencache "github.com/d-kv/backend-travel-app/pkg/infra/cache/token"
)

func (c *Controller) Authenticate(ctx context.Context, refreshToken string) (string, error) {
	const mName = "Controller.Authenticate"

	uID, err := c.tokenCache.UserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, tokencache.ErrRefreshTokenNotFound) {
			log.Warn().
				Str("method", mName).
				Err(err).
				Msg("error from tokenCache")

			return "", icontrollerv0.ErrRefreshTokenNotFound
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from tokenCache")

		return "", err
	}
	return uID, nil
}
