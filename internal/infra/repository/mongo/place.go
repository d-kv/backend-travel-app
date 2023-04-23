package mongo

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("unable to create %s index", name)
	}

	log.Info().
		Msgf("%s index created", name)

	return &PlaceStore{
		coll: coll,
	}
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
		log.Warn().
			Err(err)
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
		log.Warn().
			Err(err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Info().
			Err(irepository.ErrPlaceNotFound)
		return irepository.ErrPlaceNotFound
	}

	if res.DeletedCount > 1 {
		log.Error().
			Err(irepository.ErrUUIDDuplicate)
		return irepository.ErrUUIDDuplicate
	}
	return nil
}

// Place returns place with given UUID.
func (p *PlaceStore) Place(ctx context.Context, uuid string) (*place.Place, error) {
	res := p.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrPlaceNotFound
	}

	if err != nil {
		log.Warn().
			Err(err)
		return nil, err
	}

	var place *place.Place
	err = res.Decode(&place)
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return place, nil
}

// Places returns places.
func (p *PlaceStore) Places(ctx context.Context, skipN int64, resN int64) ([]place.Place, error) {
	opts := options.
		Find().
		SetLimit(resN).
		SetSkip(skipN)

	cursor, err := p.coll.Find(ctx, bson.D{}, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrUserNotFound
	}
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return places, nil
}

// PlacesByCategory returns places with given category.
func (p *PlaceStore) PlacesByCategory(ctx context.Context,
	mCtgs []category.MainCategory, sCtgs []category.SubCategory, skipN int64, resN int64) ([]place.Place, error) {
	if mCtgs == nil {
		mCtgs = []category.MainCategory{}
	}

	if sCtgs == nil {
		sCtgs = []category.SubCategory{}
	}

	opts := options.
		Find().
		SetLimit(resN).
		SetSkip(skipN)

	filter := bson.D{{
		Key: "$or", Value: []interface{}{
			bson.M{"category.main": bson.D{{Key: "$in", Value: mCtgs}}},
			bson.M{"category.sub": bson.D{{Key: "$in", Value: sCtgs}}},
		}}}

	cursor, err := p.coll.Find(ctx, filter, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrPlaceNotFound
	}
	if err != nil {
		log.Warn().
			Err(err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return places, nil
}

// PlacesByDistance returns places from nearest to farthest.
func (p *PlaceStore) PlacesByDistance(ctx context.Context, geoQ *query.Geo, skipN int64, resN int64) ([]place.Place, error) {
	gCenterJSON := bson.M{
		"type":        "Point",
		"coordinates": []float64{geoQ.Center.Longitude, geoQ.Center.Latitude},
	}

	opts := options.
		Find().
		SetLimit(resN).
		SetSkip(skipN)

	cursor, err := p.coll.Find(ctx, bson.M{
		"location.geo": bson.M{
			"$near": bson.M{
				"$geometry":    gCenterJSON,
				"$minDistance": geoQ.Min,
				"$maxDistance": geoQ.Max,
			},
		},
	}, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrPlaceNotFound
	}
	if err != nil {
		log.Warn().
			Err(err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return places, nil
}
