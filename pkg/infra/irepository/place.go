package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
)

type PlaceI interface {
	GetAll(ctx context.Context) ([]place.Place, error)
	Create(ctx context.Context, place *place.Place) error
	Delete(ctx context.Context, id string) error

	Get(ctx context.Context, id string) (*place.Place, error)
	GetByCategory(ctx context.Context, category category.Category) ([]place.Place, error)
	GetNearby(ctx context.Context, getQuery query.Geo) ([]place.Place, error)
}
