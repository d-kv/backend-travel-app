package mongo

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

const IndexCreationTimeout = 10

// PlaceStore with CRUD-like operations on the Place object.
type PlaceStore struct {
	coll *mongo.Collection
}

var _ irepository.PlaceI = (*PlaceStore)(nil)

// NewPlaceStore is a default ctor.
func NewPlaceStore(coll *mongo.Collection) *PlaceStore {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "location.geo", Value: "2dsphere"},
		},
	}
	name, err := coll.
		Indexes().
		CreateOne(
			context.Background(),
			indexModel,
		)

	log.Info().Msgf("NewPlaceStore: Index building done: %s", name)
	if err != nil {
		log.Panic().Msgf("NewPlaceStore: unable to create %s index", name)
	}

	return &PlaceStore{
		coll: coll,
	}
}

// GetAll returns all places.
func (p *PlaceStore) GetAll(ctx context.Context) ([]place.Place, error) {
	cursor, err := p.coll.Find(ctx, bson.D{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().Msgf("UserStore.GetByID: %v", err)
		return nil, irepository.ErrUserNotFound
	}
	if err != nil {
		log.Error().Msgf("UserStore.GetAll: %v", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().Msgf("PlaceStore.GetAll: %v", err)
		return nil, err
	}

	return places, nil
}

// Create creates a new place.
//
// UUID field must be populated.
func (p *PlaceStore) Create(ctx context.Context, place *place.Place) error {
	if place.UUID == "" {
		return irepository.ErrUUIDNotPopulated
	}

	_, err := p.coll.InsertOne(ctx, place)
	if err != nil {
		log.Warn().Msgf("PlaceStore.Create: %v", err)
		return err
	}

	return nil
}

// Delete deletes place with given UUID.
func (p *PlaceStore) Delete(ctx context.Context, uuid string) error {
	res, err := p.coll.DeleteOne(ctx, bson.M{
		"_id": uuid,
	})
	if err != nil {
		log.Warn().Msgf("PlaceStore.Delete: %v", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Warn().Msgf("PlaceStore.Delete: %v", irepository.ErrPlaceNotFound)
		return irepository.ErrPlaceNotFound
	}

	if res.DeletedCount > 1 {
		log.Error().Msgf("PlaceStore.Delete: %v", irepository.ErrUUIDDuplicate)
		return irepository.ErrUUIDDuplicate
	}
	return nil
}

// Get returns place with given UUID.
func (p *PlaceStore) Get(ctx context.Context, uuid string) (*place.Place, error) {
	res := p.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().Msgf("PlaceStore.Get: %v", err)
		return nil, irepository.ErrPlaceNotFound
	}

	if err != nil {
		log.Warn().Msgf("PlaceStore.Get: %v", err)
		return nil, err
	}

	var place *place.Place
	err = res.Decode(&place)
	if err != nil {
		log.Error().Msgf("PlaceStore.Get: %v", err)
		return nil, err
	}

	return place, nil
}

// GetByCategory returns places with given category.
func (p *PlaceStore) GetByCategory(ctx context.Context,
	mCtgs []category.MainCategory, sCtgs []category.SubCategory) ([]place.Place, error) {
	if mCtgs == nil {
		mCtgs = []category.MainCategory{}
	}

	if sCtgs == nil {
		sCtgs = []category.SubCategory{}
	}

	filter := bson.D{{
		Key: "$or", Value: []interface{}{
			bson.M{"category.main": bson.D{{Key: "$in", Value: mCtgs}}},
			bson.M{"category.sub": bson.D{{Key: "$in", Value: sCtgs}}},
		}}}

	cursor, err := p.coll.Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().Msgf("PlaceStore.GetByCategory: %v", err)
		return nil, irepository.ErrPlaceNotFound
	}
	if err != nil {
		log.Warn().Msgf("PlaceStore.GetByCategory: %v", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().Msgf("PlaceStore.GetByCategory: %v", err)
		return nil, err
	}

	return places, nil
}

// GetNearby returns places from nearest to farthest.
func (p *PlaceStore) GetNearby(ctx context.Context, geoQ *query.Geo) ([]place.Place, error) {
	gCenterJSON := bson.M{
		"type":        "Point",
		"coordinates": []float64{geoQ.Center.Longitude, geoQ.Center.Latitude},
	}

	cursor, err := p.coll.Find(ctx, bson.M{
		"location.geo": bson.M{
			"$near": bson.M{
				"$geometry":    gCenterJSON,
				"$minDistance": geoQ.Min,
				"$maxDistance": geoQ.Max,
			},
		},
	})
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().Msgf("PlaceStore.GetNearby: %v", err)
		return nil, irepository.ErrPlaceNotFound
	}
	if err != nil {
		log.Warn().Msgf("PlaceStore.GetNearby: %v", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().Msgf("PlaceStore.GetNearby: %v", err)
		return nil, err
	}

	return places, nil
}
