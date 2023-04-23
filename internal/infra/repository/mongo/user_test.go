package mongo //nolint:testpackage // Need internals of repository

import (
	"context"
	"testing"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals // Using global var in tests
var uStore *UserStore

const UserColl = "Users"

func initEmptyUserStore() {
	coll := dbClient.
		Database(mongoDB).
		Collection(UserColl)

	_ = coll.Database().Drop(context.Background())

	uStore = NewUserStore(coll)
}

func TestUserCreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyUserStore()
	assert := assert.New(t)

	id := uuid.NewString()

	wantU := user.New(
		user.WithUUID(id),
		user.WithPremium(true),
	)

	err := uStore.Create(context.Background(), wantU)
	assert.NoError(err,
		"must create user without errors")
	gotU, err := uStore.User(context.Background(), id)
	assert.NoError(err,
		"must return user without errors")

	assert.Equal(wantU, gotU)
}

func TestUserUpdateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyUserStore()
	assert := assert.New(t)

	id := uuid.NewString()

	wantU := user.New(
		user.WithUUID(id),
		user.WithPremium(true),
	)

	err := uStore.Create(context.Background(), wantU)
	assert.NoError(err,
		"must create user without errors")
	gotU, err := uStore.User(context.Background(), id)
	assert.NoError(err,
		"must return user without errors")

	assert.Equal(wantU, gotU)

	wantU.Premium = false

	err = uStore.Update(context.Background(), id, wantU)
	assert.NoError(err,
		"must update without any error")
	gotU, err = uStore.User(context.Background(), id)
	assert.NoError(err,
		"must return user without errors")
	assert.Equal(wantU, gotU,
		"must be new user with premium disabled")
}
