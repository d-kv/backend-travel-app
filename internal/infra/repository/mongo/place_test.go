package mongo //nolint:testpackage // Need internals of repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

//nolint:gochecknoglobals // Using global var in tests
var plStore *PlaceStore

const mongoURI = "mongodb://localhost:27017"
const mongoDB = "afterwork_test"
const mongoCollName = "Places"

func initEmptyPlaceStore() {
	cl, err := NewClient(mongoURI, 3*time.Second)
	if err != nil {
		panic(fmt.Sprintf("initEmptyPlaceStore: %v", err))
	}
	coll := cl.
		Database(mongoDB).
		Collection(mongoCollName)

	_ = coll.Database().Drop(context.Background())

	plStore = NewPlaceStore(coll)
}

func TestCreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
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
	initEmptyPlaceStore()
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
	initEmptyPlaceStore()
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
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create without any error")

	psExpected, err := plStore.GetAll(context.Background())

	assert.NoError(err,
		"must return all places without any error")

	ps := []place.Place{*p1, *p2}

	assert.Equal(ps, psExpected)
}

func TestGetByCategoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("My culture place #1"),
		place.WithCategory(category.NewCulture(category.CC_GALLERY)),
	)

	p2 := place.New(
		place.WithAddress("Street 2B"),
		place.WithName("My culture place #2"),
		place.WithCategory(category.NewCulture(category.CC_LIBRARY)),
	)

	p3 := place.New(
		place.WithAddress("Street 2C"),
		place.WithName("My food Place"),
		place.WithCategory(category.NewFood(category.FC_BAR)),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p3),
		"must create without any error")

	cultPlGot, err := plStore.GetByCategory(context.Background(), category.NewCulture(category.CC_UNSPECIFIED))

	assert.NoError(err,
		"must return all places without any error")

	cultPlaceWant := []place.Place{*p1, *p2}

	assert.Equal(cultPlaceWant, cultPlGot)

	foodPlGot, err := plStore.GetByCategory(context.Background(), category.NewFood(category.FC_UNSPECIFIED))

	assert.NoError(err,
		"must return all places without any error")

	foodPlaceWant := []place.Place{*p3}

	assert.Equal(foodPlaceWant, foodPlGot)
}

func TestGetNearbyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initEmptyPlaceStore()
	assert := assert.New(t)

	p1 := place.New(
		place.WithAddress("Street 2A"),
		place.WithName("My culture place #1"),
		place.WithCategory(category.NewCulture(category.CC_GALLERY)),
		place.WithLatLng(55.0, 55.0),
	)

	p2 := place.New(
		place.WithAddress("Street 2B"),
		place.WithName("My culture place #2"),
		place.WithCategory(category.NewCulture(category.CC_LIBRARY)),
		place.WithLatLng(50.0, 50.0),
	)

	p3 := place.New(
		place.WithAddress("Street 2C"),
		place.WithName("My food Place"),
		place.WithCategory(category.NewFood(category.FC_BAR)),
		place.WithLatLng(60.0, 60.0),
	)

	assert.NoError(plStore.Create(context.Background(), p1),
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p2),
		"must create without any error")

	assert.NoError(plStore.Create(context.Background(), p3),
		"must create without any error")

	geoQ := query.Geo{
		Center: util.NewLatLng(51, 51),
		Max:    10000000000,
	}
	plsGot, err := plStore.GetNearby(context.Background(), geoQ)

	assert.NoError(err,
		"must return all places without any error")

	plsWant := []place.Place{*p2, *p1, *p3}

	assert.Equal(plsWant, plsGot)
}
