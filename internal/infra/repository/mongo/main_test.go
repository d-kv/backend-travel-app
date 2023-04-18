package mongo //nolint:testpackage // Need internals of repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

//nolint:gochecknoglobals // Using global var in tests
var dbClient *mongo.Client

const mongoURI = "mongodb://localhost"
const mongoDB = "afterwork_test"
const connTimeout = 3 * time.Second

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

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
		dbClient, err = NewClient(
			fmt.Sprintf("%s:%s", mongoURI, resource.GetPort("27017/tcp")),
			connTimeout,
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
