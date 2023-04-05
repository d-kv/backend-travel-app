package igateway

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type PlaceProviderI interface {
	Get(ctx context.Context, category category.Category, ll *util.LatLng, resultN uint32, skipN uint32) ([]place.Place, error)
}
