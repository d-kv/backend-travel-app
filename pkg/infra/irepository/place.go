package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
)

type PlaceI interface {
	Places(ctx context.Context) ([]place.Place, error)
	Create(ctx context.Context, place *place.Place) error
	Delete(ctx context.Context, id string) error

	Place(ctx context.Context, id string) (*place.Place, error)
	PlacesByCategory(ctx context.Context,
		mCtgs []category.MainCategory, sCtgs []category.SubCategory) ([]place.Place, error)
	PlacesByDistance(ctx context.Context, getQuery *query.Geo) ([]place.Place, error)
}
