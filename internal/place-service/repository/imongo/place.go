package imongo

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
	"github.com/d-kv/backend-travel-app/pkg/place-service/repository"
)

const IndexCreationTimeout = 10

// PlaceStore with CRUD-like operations on the Place object.
type PlaceStore struct {
	coll *mongo.Collection
}

var _ repository.PlaceProvider = (*PlaceStore)(nil)

// New is a default ctor.
func New(coll *mongo.Collection) *PlaceStore {
	const mName = "NewPlaceStore"

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
			Str("method", mName).
			Err(err).
			Msgf("unable to create %s index", name)
	}

	log.Info().
		Str("method", mName).
		Msgf("%s index created", name)

	return &PlaceStore{
		coll: coll,
	}
}

// Create creates a new place.
//
// UUID field must be populated.
func (p *PlaceStore) Create(ctx context.Context, place *model.Place) error {
	const mName = "PlaceStore.Create"

	if place.UUID == "" {
		return repository.ErrUUIDNotPopulated
	}

	_, err := p.coll.InsertOne(ctx, place)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return err
	}

	return nil
}

// Delete deletes place with given UUID.
func (p *PlaceStore) Delete(ctx context.Context, uuid string) error {
	const mName = "PlaceStore.Delete"

	res, err := p.coll.DeleteOne(ctx, bson.M{
		"_id": uuid,
	})
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return err
	}

	if res.DeletedCount == 0 {
		log.Info().
			Str("method", mName).
			Str("uuid", uuid).
			Msg("no places with given UUID")

		return repository.ErrPlaceNotFound
	}
	return nil
}

// Place returns place with given UUID.
func (p *PlaceStore) Place(ctx context.Context, uuid string) (*model.Place, error) {
	const mName = "PlaceStore.Place"

	res := p.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Str("method", mName).
			Str("uuid", uuid).
			Err(err).
			Msg("no places with given UUID")

		return nil, repository.ErrPlaceNotFound
	}

	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return nil, err
	}

	var place *model.Place
	err = res.Decode(&place)
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error while decoding")

		return nil, err
	}

	return place, nil
}

// Places returns places.
func (p *PlaceStore) Places(ctx context.Context, skipN int64, resN int64) ([]model.Place, error) {
	const mName = "PlaceStore.Places"

	opts := options.
		Find().
		SetLimit(resN).
		SetSkip(skipN)

	cursor, err := p.coll.Find(ctx, bson.D{}, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, repository.ErrPlaceNotFound
	}
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return nil, err
	}

	var places []model.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error while decoding")

		return nil, err
	}

	return places, nil
}

// PlacesByCategory returns places with given category.
func (p *PlaceStore) PlacesByCategory(ctx context.Context, mainCtgs []category.Main, subCtgs []category.Sub, skipN int64, resN int64) ([]model.Place, error) {
	const mName = "PlaceStore.PlacesByCategory"

	unsortedPlaces, err := p.Places(ctx, skipN, resN)
	if err != nil {
		return nil, err
	}

	var places []model.Place

	for _, pl := range unsortedPlaces {
		for _, cat := range mainCtgs {
			if slices.Contains(pl.Category.Main, cat) {
				places = append(places, pl)
			}
		}

		for _, cat := range subCtgs {
			if slices.Contains(pl.Category.Sub, cat) {
				places = append(places, pl)
			}
		}
	}
	return places, nil
}

// PlacesByDistance returns places from nearest to farthest.
func (p *PlaceStore) PlacesByDistance(ctx context.Context,
	geoQ *util.GeoToken, skipN int64, resN int64) ([]model.Place, error) {
	const mName = "PlaceStore.PlacesByDistance"

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
			Str("method", mName).
			Err(err).
			Msg("no places for the given criteria")

		return nil, repository.ErrPlaceNotFound
	}
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return nil, err
	}

	var places []model.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error while decoding")

		return nil, err
	}

	return places, nil
}
