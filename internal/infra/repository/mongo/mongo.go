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
	defer cancel()

	l.Info("NewClient: mongoDB uri:", uri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		l.Error("NewClient:, %v", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		l.Error("NewClient: %v", err)
	}

	l.Info("NewClient: Connected to MongoDB")
	return client, nil
}
