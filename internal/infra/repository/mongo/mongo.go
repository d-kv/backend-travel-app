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

	log.Println("Connect: attempt to connect to mongoDB")

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Printf("Connect: mongoDB connection error:, %s\n", err)
		return nil, err
	}

	log.Println("Connect: mongoDB successful connection")
	return client, nil
}
