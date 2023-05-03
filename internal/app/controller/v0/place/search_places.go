package iplace_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	place_controller_v0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	placerepo "github.com/d-kv/backend-travel-app/pkg/infra/repository/place"
)

func (c *PlaceController) SearchPlaces(ctx context.Context, geoQ *query.Geo,
	_ []category.MainCategory, _ []category.SubCategory, skipN int64, resN int64) ([]place.Place, error) {
	const mName = "PlaceController.SearchPlaces"

	places, err := c.placeProvider.PlacesByDistance(ctx, geoQ, skipN, resN)
	if err != nil {
		if errors.Is(err, placerepo.ErrPlaceNotFound) {
			log.Info().
				Str("method", mName).
				Err(err).
				Msg("no places for the given criteria")

			return nil, place_controller_v0.ErrNoPlaces
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from placeProvider")

		return nil, err
	}

	return places, nil
}
