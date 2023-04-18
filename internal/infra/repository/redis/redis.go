package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewClient(uri string, connTimeout time.Duration) (*redis.Client, error) {
	log.Info().Msgf("redis.NewClient: redis uri: %s", uri)
	cl := redis.NewClient(
		&redis.Options{
			Addr: uri,
		},
	)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	err := cl.Ping(ctx).Err()
	if err != nil {
		log.Error().Msgf("NewClient: %v", err)
		return nil, err
	}

	log.Info().Msgf("redis.NewClient: redis uri: %s", uri)

	return cl, nil
}
