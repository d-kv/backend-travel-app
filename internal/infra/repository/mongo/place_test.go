package mongo //nolint:testpackage // Need internals of repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

//nolint:gochecknoglobals // Using global var in tests
var plStore *PlaceStore

const mongoURI = "mongodb://localhost:27017"
const mongoDB = "Afterwork-DB-Test"
const mongoCollName = "Places"

//nolint:gochecknoinits // Using init() in tests
func init() {
	cl, err := NewClient(mongoURI)
	if err != nil {
		panic(fmt.Errorf("init: %w", err))
	}

	plStore = NewPlaceStore(cl.
		Database(mongoDB).
		Collection(mongoCollName),
	)

	docN, err := plStore.coll.CountDocuments(context.Background(), bson.D{})

	if err != nil {
		panic(fmt.Errorf("init: %w", err))
	}
	if docN != 0 {
		panic(fmt.Errorf("init: collection is not empty"))
		// Make sure that it is safe to use the collection in tests & erase it manually
	}
}

func dropPlaceStore() {
	err := plStore.coll.Drop(context.Background())
	if err != nil {
		panic(fmt.Errorf("dropPlaceStore: %w", err))
	}
}

func TestCreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	dropPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithUUID(uuid.New().String()),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create without any error")

	duplID := uuid.New().String()

	p2 := place.New(
		place.WithUUID(duplID),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create without any error")

	p3 := place.New(
		place.WithUUID(duplID),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.True(mongo.IsDuplicateKeyError(plStore.Create(context.Background(), p3)),
		"must be primitive.ObjectID duplicate error")

	p2Expected, err := plStore.Get(context.Background(), duplID)

	assert.NoError(err, "must return place without error")
	assert.Equal(p2, p2Expected, "must be the same place which was inserted")
}

func TestDeleteIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	dropPlaceStore()
	assert := assert.New(t)

	id := uuid.New().String()

	p := place.New(
		place.WithUUID(id),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p),
		"must create without any error")

	assert.NoError(plStore.Delete(context.Background(), id),
		"must delete without any error")

	assert.ErrorIs(plStore.Delete(context.Background(), id), irepository.ErrPlaceNotFound,
		"must be irepository.ErrPlaceNotFound")
}

func TestGetByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	dropPlaceStore()
	assert := assert.New(t)

	id := uuid.New().String()

	p := place.New(
		place.WithUUID(id),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p),
		"must create without any error")

	pExpected, err := plStore.Get(context.Background(), id)

	assert.NoError(err,
		"must return place without any error")

	assert.Equal(p, pExpected)
}

func TestGetAllIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	dropPlaceStore()
	assert := assert.New(t)

	id1 := uuid.New().String()
	id2 := uuid.New().String()

	p1 := place.New(
		place.WithUUID(id1),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace1"),
	)

	p2 := place.New(
		place.WithUUID(id2),
		place.WithAddress("Street 2B"),
		place.WithName("MyPlace2"),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create without any error")

	psExpected, err := plStore.GetAll(context.Background())

	assert.NoError(err,
		"must return all places without any error")

	ps := []*place.Place{p1, p2}

	assert.Equal(ps, psExpected)
}
