package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/d-kv/backend-travel-app/pkg/infra/ilogger"
)

func NewClient(l ilogger.LoggerI, uri string, connTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)

	l.Info("NewClient: attempt to connect to mongoDB at:", uri)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		l.Error("NewClient: mongoDB connection error:, %v\n", err)
		return nil, err
	}

	l.Info("NewClient: mongoDB successful connection")
	return client, nil
}
