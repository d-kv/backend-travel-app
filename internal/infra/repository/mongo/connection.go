package mongo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(uri string, connTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	log.Info().
		Str("mongoDB uri", uri).
		Msg("trying to connect to mongoDB")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	log.Info().
		Msg("connected to mongoDB")
	return client, nil
}
