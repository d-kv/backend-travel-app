package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
)

type PlaceI interface {
	GetAll(context.Context) ([]place.Place, error)
	Create(context.Context, *place.Place) error
	Delete(context.Context, string) error

	Get(context.Context, string) (*place.Place, error)
	GetByCategory(context.Context, category.Category) ([]place.Place, error)
	GetNearby(context.Context, query.Geo) ([]place.Place, error)
}
