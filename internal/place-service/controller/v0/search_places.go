package ictrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	ctrl_v0 "github.com/d-kv/backend-travel-app/pkg/place-service/controller/v0"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
	"github.com/d-kv/backend-travel-app/pkg/place-service/repository"
)

func (c *PlaceController) SearchPlaces(ctx context.Context, geoQ *util.GeoToken, category *category.Category,
	skipN int64, resN int64) ([]model.Place, error) {
	const mName = "PlaceController.SearchPlaces"

	places, err := c.placeProvider.PlacesByDistance(ctx, geoQ, skipN, resN)
	if err != nil {
		if errors.Is(err, repository.ErrPlaceNotFound) {
			return nil, ctrl_v0.ErrNoPlaces
		}

		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from placeProvider")

		return nil, err
	}

	suitablePlaces := filterByCategory(places, category)

	return suitablePlaces, nil
}
