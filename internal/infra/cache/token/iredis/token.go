package redistoken

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	tokencache "github.com/d-kv/backend-travel-app/pkg/infra/cache/token"
)

type TokenCache struct {
	db *redis.Client
}

var _ tokencache.TokenCacher = (*TokenCache)(nil)

func NewTokenCache(cl *redis.Client) *TokenCache {
	return &TokenCache{
		db: cl,
	}
}

func (t *TokenCache) SetUserID(ctx context.Context, rToken, userUUID string) error {
	err := t.db.Set(ctx, rToken, userUUID, 0).Err()
	if err != nil {
		log.Error().
			Err(err)
		return err
	}
	return nil
}

func (t *TokenCache) UserID(ctx context.Context, rToken string) (string, error) {
	uID, err := t.db.Get(ctx, rToken).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Info().
				Err(err)
			return "", tokencache.ErrRefreshTokenNotFound
		}

		log.Error().
			Err(err)
		return "", err
	}
	return uID, nil
}
