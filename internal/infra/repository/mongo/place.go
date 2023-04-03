package mongo

import (
	"context"
	"errors"
	"log"

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
	name, err := coll.Indexes().
		CreateOne(context.Background(),
			indexModel)
	log.Println(name)
	if err != nil {
		panic("NewPlaceStore: unable to create geoIndex")
	}
	return &PlaceStore{
		coll: coll,
	}
}

// GetAll returns all places.
func (p *PlaceStore) GetAll(ctx context.Context) ([]place.Place, error) {
	cursor, err := p.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("PlaceStore.GetAll: db error: %v\n", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Printf("PlaceStore.GetAll: decoding error: %v\n", err)
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
		log.Printf("PlaceStore.Create: DB error: %v\n", err)
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
		log.Printf("PlaceStore.Delete: db error: %v\n", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Printf("PlaceStore.Delete: db error: %v\n", irepository.ErrPlaceNotFound)
		return irepository.ErrPlaceNotFound
	}

	if res.DeletedCount > 1 {
		log.Printf("PlaceStore.Delete: db error: %v\n", irepository.ErrUUIDDuplicate)
		return irepository.ErrUUIDDuplicate
	}

	return nil
}

// GetByID returns place with given UUID.
func (p *PlaceStore) Get(ctx context.Context, uuid string) (*place.Place, error) {
	res := p.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Printf("PlaceStore.Get: db error: %v\n", err)
		return nil, irepository.ErrPlaceNotFound
	}

	if err != nil {
		log.Printf("PlaceStore.Get: db error: %v\n", err)
		return nil, err
	}

	var place *place.Place
	err = res.Decode(&place)
	if err != nil {
		log.Printf("PlaceStore.Get: decoding error: %v\n", err)
		return nil, err
	}

	return place, nil
}

// GetByCategory returns places with given category.
func (p *PlaceStore) GetByCategory(ctx context.Context, category category.Category) ([]place.Place, error) {
	cursor, err := p.coll.Find(ctx, bson.M{
		"category.main_category": category.MainCategoryString(), // TODO: add aggregation by subCategory
	})
	if err != nil {
		log.Printf("PlaceStore.GetByCategory: db error: %v\n", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Printf("PlaceStore.GetByCategory: decoding error: %v\n", err)
		return nil, err
	}

	return places, nil
}

func (p *PlaceStore) GetNearby(ctx context.Context, geoQ query.Geo) ([]place.Place, error) {
	gCenterJSON := bson.M{
		"type":        "Point",
		"coordinates": []float64{geoQ.Center.Longitude, geoQ.Center.Latitude},
	}

	cursor, err := p.coll.Find(ctx, bson.M{
		"location.geo": bson.M{
			"$near": bson.M{
				"$geometry": gCenterJSON,
				// TODO: receive from request
				// "$minDistance": geoQ.Min,
				// "$maxDistance": geoQ.Max,
			},
		},
	})
	if err != nil {
		log.Printf("PlaceStore.GetNearby: db error: %v\n", err)
		return nil, err
	}

	var places []place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Printf("PlaceStore.GetNearby: decoding error: %v\n", err)
		return nil, err
	}

	return places, nil
}
