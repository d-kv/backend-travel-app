package mongo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

const mongoURI = "mongodb://localhost"
const mongoDB = "afterwork_test"

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Error().Msgf("Could not construct pool: %v", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Error().Msgf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("mongo", "6.0.5", nil)
	if err != nil {
		log.Error().Msgf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("%s:%s", mongoURI, resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})

	if err != nil {
		log.Error().Msgf("Could not connect to docker: %s", err)
	}

	// run tests
	code := m.Run()

	if err = dbClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	if err = pool.Purge(resource); err != nil {
		log.Error().Msgf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}