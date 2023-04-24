package mongoplace //nolint:testpackage // Need internals of repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	placerepo "github.com/d-kv/backend-travel-app/pkg/infra/repository/place"
)

//nolint:gochecknoglobals // Using global var in tests
var plStore *PlaceStore

const PlaceColl = "Places"

func initEmptyPlaceStore() {
	coll := dbClient.
		Database(mongoDB).
		Collection(PlaceColl)

	_ = coll.Database().Drop(context.Background())

	plStore = NewPlaceStore(coll)
}

func TestPlaceCreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
		place.WithMainCategories(category.MC_CULTURE, category.MC_HOSPITALITY),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	duplID := uuid.New().String()

	p2 := place.New(
		place.WithUUID(duplID),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
		place.WithSubCategories(category.SC_CINEMA, category.SC_BAR),
	)

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	p3 := place.New(
		place.WithUUID(duplID),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
		place.WithMainCategories(category.MC_FOOD),
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

	p := place.New(
		place.WithUUID(id),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
	)

	assert.NoError(plStore.Create(context.Background(), p),
		"must create place without errors")

	assert.NoError(plStore.Delete(context.Background(), id),
		"must delete place without errors")

	assert.ErrorIs(plStore.Delete(context.Background(), id), placerepo.ErrPlaceNotFound,
		"must be %v", placerepo.ErrPlaceNotFound)
}

func TestPlacePlaceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	id := uuid.New().String()

	p := place.New(
		place.WithUUID(id),
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace"),
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

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("MyPlace1"),
	)

	p2 := place.New(
		place.WithAddress("Street 2B"),
		place.WithName("MyPlace2"),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	psExpected, err := plStore.Places(context.Background(), 0, 0)

	assert.NoError(err,
		"must return all places without errors")

	ps := []place.Place{*p1, *p2}

	assert.Equal(ps, psExpected)
}

func TestPlacePlacesByCategoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("My culture place #1"),
		place.WithMainCategories(
			category.MC_CULTURE,
			category.MC_HOSPITALITY,
		),
	)

	p2 := place.New(
		place.WithAddress("Street 2B"),
		place.WithName("My culture place #2"),
		place.WithMainCategories(category.MC_CULTURE),
	)

	p3 := place.New(
		place.WithAddress("Street 2C"),
		place.WithName("My food Place"),
		place.WithMainCategories(category.MC_FOOD),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p3),
		"must create place without errors")

	cultPlGot, err := plStore.PlacesByCategory(
		context.Background(),
		[]category.MainCategory{category.MC_CULTURE},
		nil,
		0,
		0,
	)

	assert.NoError(err,
		"must return all places without errors")

	cultPlaceWant := []place.Place{*p1, *p2}

	assert.Equal(cultPlaceWant, cultPlGot)

	foodPlGot, err := plStore.PlacesByCategory(
		context.Background(),
		[]category.MainCategory{category.MC_FOOD},
		nil,
		0,
		0,
	)

	assert.NoError(err,
		"must return all places without errors")

	foodPlaceWant := []place.Place{*p3}

	assert.Equal(foodPlaceWant, foodPlGot)
}

func TestPlacePlacesByDistanceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("My culture place #1"),
		place.WithLatLng(55.0, 55.0),
	)

	p2 := place.New(
		place.WithAddress("Street 2B"),
		place.WithName("My culture place #2"),
		place.WithLatLng(50.0, 50.0),
	)

	p3 := place.New(
		place.WithAddress("Street 2C"),
		place.WithName("My food Place"),
		place.WithLatLng(60.0, 60.0),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create place without errors")

	assert.NoError(plStore.Create(context.Background(), p3),
		"must create place without errors")

	geoQ := &query.Geo{
		Center: util.NewLatLng(51, 51),
		Max:    10000000000,
	}
	plsGot, err := plStore.PlacesByDistance(context.Background(), geoQ, 0, 0)

	assert.NoError(err,
		"must return all places without errors")

	plsWant := []place.Place{*p2, *p1, *p3}

	assert.Equal(plsWant, plsGot)
}
