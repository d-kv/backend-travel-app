package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// PlaceStore with CRUD-like operations on the Place object.
type PlaceStore struct {
	coll *mongo.Collection
}

var _ irepository.PlaceI = (*PlaceStore)(nil)

// NewPlaceStore is a default ctor.
func NewPlaceStore(coll *mongo.Collection) *PlaceStore {
	return &PlaceStore{
		coll: coll,
	}
}

// GetAll returns all places.
func (p *PlaceStore) GetAll(ctx context.Context) ([]*place.Place, error) {
	cursor, err := p.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("PlaceStore.GetAll: db error: %s\n", err)
		return nil, err
	}

	var places []*place.Place
	err = cursor.All(ctx, &places) // FIXME: may be an overflow
	if err != nil {
		log.Printf("PlaceStore.GetAll: decoding error: %s\n", err)
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
		log.Printf("PlaceStore.Create: DB error: %s\n", err)
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
		log.Printf("PlaceStore.Delete: db error: %s\n", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Printf("PlaceStore.Delete: db error: %s\n", irepository.ErrPlaceNotFound)
		return irepository.ErrPlaceNotFound
	}

	if res.DeletedCount > 1 {
		log.Printf("PlaceStore.Delete: db error: %s\n", irepository.ErrUUIDDuplicate)
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
		log.Printf("PlaceStore.Get: db error: %s\n", err)
		return nil, irepository.ErrPlaceNotFound
	}

	if err != nil {
		log.Printf("PlaceStore.Get: db error: %s\n", err)
		return nil, err
	}

	var place *place.Place
	err = res.Decode(&place)
	if err != nil {
		log.Printf("PlaceStore.Get: decoding error: %s\n", err)
		return nil, err
	}

	return place, nil
}
