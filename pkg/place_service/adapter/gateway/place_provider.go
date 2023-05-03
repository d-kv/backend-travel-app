package gateway

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model"
	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model/util"
)

type PlaceProviderI interface {
	Get(ctx context.Context, category category.Category,
		ll *util.LatLng, resultN uint32, skipN uint32) ([]model.Place, error)
}
