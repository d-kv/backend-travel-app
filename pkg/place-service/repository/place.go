package repository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

type PlaceProvider interface {
	Create(ctx context.Context, place *model.Place) error
	Delete(ctx context.Context, id string) error
	Place(ctx context.Context, id string) (*model.Place, error)

	Places(ctx context.Context, skipN int64, resN int64) ([]model.Place, error)
	PlacesByCategory(ctx context.Context,
		mCtgs []category.Main, sCtgs []category.Sub, skipN int64, resN int64) ([]model.Place, error)
	PlacesByDistance(ctx context.Context, getQuery *util.GeoToken, skipN int64, resN int64) ([]model.Place, error)
}
