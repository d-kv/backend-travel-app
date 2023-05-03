// TODO: add tests
package iplace_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
)

//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=PlaceProvider --output=mock --case=underscore --disable-version-string --outpkg=mock

type (
	PlaceProvider interface {
		PlacesByDistance(ctx context.Context, geoQ *query.Geo, skipN int64, resN int64) ([]place.Place, error)
	}
)

type PlaceController struct {
	placeProvider PlaceProvider
}

// New is a default ctor for PlaceController.
func New(placeP PlaceProvider) *PlaceController {
	return &PlaceController{
		placeProvider: placeP,
	}
}
