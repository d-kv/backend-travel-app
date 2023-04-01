// TODO: add tests
package controller

import (
	"context"
	"log"

	"github.com/d-kv/backend-travel-app/pkg/app/icontroller"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// Controller defines a place service controller.
type Controller struct {
	placeStore irepository.PlaceI
	userStore  irepository.UserI
}

var _ icontroller.ControllerI = (*Controller)(nil)

// New is a default ctor for Controller.
func New(pStore irepository.PlaceI, uStore irepository.UserI) *Controller {
	return &Controller{
		placeStore: pStore,
		userStore:  uStore,
	}
}

// TODO: move identity check to adapter layer using interceptors
func (c *Controller) GetAchievements(ctx context.Context, userUUID string) (*util.Achievements, error) {
	u, err := c.userStore.GetByID(ctx, userUUID)
	if err != nil {
		log.Printf("Controller.Authorize: db error: %s\n", err)
		return nil, err
	}

	return &u.Achievements, nil
}

// TODO: move identity check to adapter layer using interceptors
func (c *Controller) GetPlaces(ctx context.Context, ctg category.Category, gCenter util.LatLng) ([]place.Place, error) {
	if !ctg.MainCategoryIsSpecified() {
		return nil, icontroller.ErrCategoryNotSpecified
	}

	places, err := c.placeStore.GetByCategory(ctx, ctg)
	if err != nil {
		log.Printf("Controller.GetPlaces: db error: %s\n", err)
		return nil, err
	}

	return places, nil
}
