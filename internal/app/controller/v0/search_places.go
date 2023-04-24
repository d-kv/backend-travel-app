package controllerv0

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	icontrollerv0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	placerepo "github.com/d-kv/backend-travel-app/pkg/infra/repository/place"
)

func (c *Controller) SearchPlaces(ctx context.Context, geoQ *query.Geo,
	_ []category.MainCategory, _ []category.SubCategory, skipN int64, resN int64) ([]place.Place, error) {
	places, err := c.placeProvider.PlacesByDistance(ctx, geoQ, skipN, resN)
	if err != nil {
		if errors.Is(err, placerepo.ErrPlaceNotFound) {
			log.Info().
				Err(err)
			return nil, icontrollerv0.ErrNoPlaces
		}

		log.Error().
			Err(err)
		return nil, err
	}

	return places, nil
}
