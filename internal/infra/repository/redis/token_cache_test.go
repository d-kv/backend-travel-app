package redis //nolint:testpackage // Need internals of repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals // Using global var in tests
var tStore *TokenCache

func init() { //nolint:gochecknoinits // Using init for tests
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func initEmptyTokenStore() {
	res, err := dbClient.FlushAll(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("initEmptyTokenStore: %v", err))
	}
	log.Info().Msgf("initEmptyTokenStore: %s", res)

	tStore = NewTokenStore(dbClient)
}

func TestSetUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyTokenStore()
	assert := assert.New(t)

	gotRToken := "MyRefreshToken"
	gotUUID := "MyUUID"
	gotUUID2 := "MyUUID2"

	assert.NoError(tStore.SetUserID(
		context.Background(), gotRToken, gotUUID),
		"must insert without any error")

	assert.NoError(tStore.SetUserID(
		context.Background(), gotRToken, gotUUID2),
		"must replace without any error")
}

func TestUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyTokenStore()
	assert := assert.New(t)

	wantRToken := "MyRefreshToken"
	wantUUID := "MyUUID"
	wantUUID2 := "MyUUID2"

	_, err := tStore.UserID(context.Background(), wantRToken)
	assert.ErrorIs(err, irepository.ErrRefreshTokenNotFound,
		"must return ErrRefreshTokenNotFound")

	assert.NoError(tStore.SetUserID(
		context.Background(), wantRToken, wantUUID),
		"must insert without any error")

	gotUUID, err := tStore.UserID(context.Background(), wantRToken)
	assert.NoError(err,
		"must return uuid without any error")
	assert.Equal(wantUUID, gotUUID,
		"must be the same")

	assert.NoError(tStore.SetUserID(
		context.Background(), wantRToken, wantUUID2),
		"must replace without any error")

	gotUUID2, err := tStore.UserID(context.Background(), wantRToken)
	assert.NoError(err,
		"must return uuid without any error")
	assert.Equal(wantUUID2, gotUUID2,
		"must be the same")
}
