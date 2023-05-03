package imongo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(uri string, connTimeout time.Duration) (*mongo.Client, error) {
	const mName = "imongo.NewClient"

	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	log.Info().
		Str("method", mName).
		Str("mongoDB uri", uri).
		Msg("trying to connect to mongoDB")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("unable to connect to mongoDB")

		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("unable to ping mongoDB")

		return nil, err
	}

	log.Info().
		Str("method", mName).
		Msg("connected to mongoDB")
	return client, nil
}
