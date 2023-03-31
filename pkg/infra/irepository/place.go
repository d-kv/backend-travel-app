package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
)

type PlaceI interface {
	GetAll(context.Context) ([]*place.Place, error)
	Create(context.Context, *place.Place) error
	Delete(context.Context, string) error

	GetByID(context.Context, string) (*place.Place, error)
}
