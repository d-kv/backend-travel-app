package controllerv0

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	icontrollerv0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0"
	tokencache "github.com/d-kv/backend-travel-app/pkg/infra/cache/token"
)

func (c *Controller) Authenticate(ctx context.Context, refreshToken string) (string, error) {
	uID, err := c.TokenCache.UserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, tokencache.ErrRefreshTokenNotFound) {
			log.Warn().
				Err(err)
			return "", icontrollerv0.ErrRefreshTokenNotFound
		}

		log.Error().
			Err(err)
		return "", err
	}
	return uID, nil
}
