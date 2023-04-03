package igateway

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type PlaceProviderI interface {
	Get(context.Context, category.Category, util.LatLng, uint32) ([]place.Place, error)
}
