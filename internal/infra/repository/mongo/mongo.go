package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connTimeout = 10

func NewClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)

	log.Println("NewClient: attempt to connect to mongoDB at:", uri)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Printf("NewClient: mongoDB connection error:, %v\n", err)
		return nil, err
	}

	log.Println("NewClient: mongoDB successful connection")
	return client, nil
}
