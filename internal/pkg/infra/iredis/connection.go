package iredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewClient(uri string, connTimeout time.Duration) (*redis.Client, error) {
	const mName = "iredis.NewClient"

	log.Info().
		Str("method", mName).
		Str("redis uri", uri).
		Msg("trying to connect to redis")

	cl := redis.NewClient(
		&redis.Options{
			Addr: uri,
		},
	)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	err := cl.Ping(ctx).Err()
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("unable to ping redis")

		return nil, err
	}

	log.Info().
		Str("method", mName).
		Msg("connected to redis")

	return cl, nil
}
