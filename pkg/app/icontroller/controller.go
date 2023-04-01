package icontroller

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type ControllerI interface {
	// TODO: move auth to interceptors & refactor Authorize into ?Bootstrap
	Authorize(context.Context, string, string, util.Version) (*util.Achievements, error)
	GetPlaces(context.Context, string, string, category.Category, util.LatLng) ([]place.Place, error)
}
