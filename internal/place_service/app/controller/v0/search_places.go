package iplace_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	controller_v0 "github.com/d-kv/backend-travel-app/pkg/place_service/app/controller/v0"
	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model"
	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place_service/domain/model/util"
	"github.com/d-kv/backend-travel-app/pkg/place_service/infra/repository"
)

func (c *PlaceController) SearchPlaces(ctx context.Context, geoQ *util.GeoQuery,
	_ []category.MainCategory, _ []category.SubCategory, skipN int64, resN int64) ([]model.Place, error) {
	const mName = "PlaceController.SearchPlaces"

	places, err := c.placeProvider.PlacesByDistance(ctx, geoQ, skipN, resN)
	if err != nil {
		if errors.Is(err, repository.ErrPlaceNotFound) {
			log.Info().
				Str("method", mName).
				Err(err).
				Msg("no places for the given criteria")

			return nil, controller_v0.ErrNoPlaces
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from placeProvider")

		return nil, err
	}

	return places, nil
}