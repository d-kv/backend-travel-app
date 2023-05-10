package imongo //nolint:testpackage // Need internals of repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
	"github.com/d-kv/backend-travel-app/pkg/place-service/repository"
)

//nolint:gochecknoglobals // Using global var in tests
var plStore *PlaceStore

const PlaceColl = "Places"

func initEmptyPlaceStore() {
	coll := dbClient.
		Database(mongoDB).
		Collection(PlaceColl)

	_ = coll.Database().Drop(context.Background())

	plStore = New(coll)
}

func TestPlaceCreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := model.New(
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace"),
		model.WithMainCategories(
			category.MC_CULTURE,
			category.MC_HOSPITALITY,
		),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	duplID := uuid.New().String()

	p2 := model.New(
		model.WithUUID(duplID),
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace"),
		model.WithSubCategories(
			category.SC_CINEMA,
			category.SC_BAR,
		),
	)

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	p3 := model.New(
		model.WithUUID(duplID),
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace"),
		model.WithMainCategories(category.MC_FOOD),
	)

	assert.True(mongo.IsDuplicateKeyError(plStore.Create(context.Background(), p3)),
		"must be primitive.ObjectID duplicate error")

	p2Expected, err := plStore.Place(context.Background(), duplID)

	assert.NoError(err,
		"must return place without errors")
	assert.Equal(p2, p2Expected,
		"must be the same place which was inserted")
}

func TestPlaceDeleteIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	id := uuid.New().String()

	p := model.New(
		model.WithUUID(id),
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p),
		"must create place without errors")

	assert.NoError(plStore.Delete(context.Background(), id),
		"must delete place without errors")

	assert.ErrorIs(plStore.Delete(context.Background(), id), repository.ErrPlaceNotFound,
		"must be %v", repository.ErrPlaceNotFound)
}

func TestPlacePlaceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	id := uuid.New().String()

	p := model.New(
		model.WithUUID(id),
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p),
		"must create place without errors")

	pExpected, err := plStore.Place(context.Background(), id)

	assert.NoError(err,
		"must return place without errors")

	assert.Equal(p, pExpected)
}

func TestPlacePlacesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := model.New(
		model.WithAddress("Street 2A"),
		model.WithName("MyPlace1"),
	)

	p2 := model.New(
		model.WithAddress("Street 2B"),
		model.WithName("MyPlace2"),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	psExpected, err := plStore.Places(context.Background(), 0, 0)

	assert.NoError(err,
		"must return all places without errors")

	ps := []model.Place{*p1, *p2}

	assert.Equal(ps, psExpected)
}

func TestPlacePlacesByDistanceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := model.New(
		model.WithAddress("Street 2A"),
		model.WithName("My culture place #1"),
		model.WithLatLng(55.0, 55.0),
	)

	p2 := model.New(
		model.WithAddress("Street 2B"),
		model.WithName("My culture place #2"),
		model.WithLatLng(50.0, 50.0),
	)

	p3 := model.New(
		model.WithAddress("Street 2C"),
		model.WithName("My food Place"),
		model.WithLatLng(60.0, 60.0),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p3),
		"must create place without errors")

	geoQ := &util.GeoToken{
		Center: util.NewLatLng(51, 51),
		Max:    10000000000,
	}
	plsGot, err := plStore.PlacesByDistance(context.Background(), geoQ, 0, 0)

	assert.NoError(err,
		"must return all places without errors")

	plsWant := []model.Place{*p2, *p1, *p3}

	assert.Equal(plsWant, plsGot)
}
