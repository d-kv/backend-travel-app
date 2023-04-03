// TODO: add tests
package controller

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/app/icontroller"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	"github.com/d-kv/backend-travel-app/pkg/infra/ilogger"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// Controller defines a place service controller.
type Controller struct {
	logger     ilogger.LoggerI
	placeStore irepository.PlaceI
	userStore  irepository.UserI
}

var _ icontroller.ControllerI = (*Controller)(nil)

// New is a default ctor for Controller.
func New(l ilogger.LoggerI, pStore irepository.PlaceI, uStore irepository.UserI) *Controller {
	return &Controller{
		logger:     l,
		placeStore: pStore,
		userStore:  uStore,
	}
}

func (c *Controller) GetAchievements(ctx context.Context, userUUID string) (*util.Achievements, error) {
	u, err := c.userStore.GetByID(ctx, userUUID)
	if err != nil {
		c.logger.Info("Controller.GetAchievements: userStore error: %v\n", err)
		return nil, err
	}

	return &u.Achievements, nil
}

func (c *Controller) GetPlaces(ctx context.Context, gCenter *util.LatLng) ([]place.Place, error) {
	geoQ := query.Geo{ // TODO: receive min & max parameters from request
		Center: gCenter,
	}

	places, err := c.placeStore.GetNearby(ctx, geoQ)
	if err != nil {
		c.logger.Info("Controller.GetPlaces: placeStore error: %v\n", err)
		return nil, err
	}

	return places, nil
}
