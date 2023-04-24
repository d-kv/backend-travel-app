package placestore

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
)

type PlaceProvider interface {
	Create(ctx context.Context, place *place.Place) error
	Delete(ctx context.Context, id string) error
	Place(ctx context.Context, id string) (*place.Place, error)

	Places(ctx context.Context, skipN int64, resN int64) ([]place.Place, error)
	PlacesByCategory(ctx context.Context,
		mCtgs []category.MainCategory, sCtgs []category.SubCategory, skipN int64, resN int64) ([]place.Place, error)
	PlacesByDistance(ctx context.Context, getQuery *query.Geo, skipN int64, resN int64) ([]place.Place, error)
}
