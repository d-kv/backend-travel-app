package gateway

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

type PlaceProviderI interface {
	Get(ctx context.Context, category category.Category,
		ll *util.LatLng, resultN uint32, skipN uint32) ([]model.Place, error)
}
