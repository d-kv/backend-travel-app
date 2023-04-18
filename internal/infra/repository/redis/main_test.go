package redis //nolint:testpackage // Need internals of repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

//nolint:gochecknoglobals // Using global var in tests
var dbClient *redis.Client

const redisURI = "localhost"
const connTimeout = 3 * time.Second

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Error().Msgf("Could not construct pool: %v", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Error().Msgf("Could not connect to Docker: %v", err)
	}

	resource, err := pool.Run("redis", "7.0.10", nil)
	if err != nil {
		log.Error().Msgf("Could not start resource: %v", err)
	}

	if err = pool.Retry(func() error {
		dbClient, err = NewClient(
			fmt.Sprintf("%s:%s", redisURI, resource.GetPort("6379/tcp")),
			connTimeout,
		)
		if err != nil {
			return err
		}

		return dbClient.Ping(context.TODO()).Err()
	}); err != nil {
		log.Error().Msgf("Could not connect to docker: %v", err)
	}
	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Error().Msgf("Could not purge resource: %v", err)
	}

	os.Exit(code)
}
