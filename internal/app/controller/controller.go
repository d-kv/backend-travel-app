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
	minVersion util.Version
}

var _ icontroller.ControllerI = (*Controller)(nil)

// New is a default ctor for Controller.
func New(pStore irepository.PlaceI, uStore irepository.UserI,
	mVersion util.Version) *Controller {
	return &Controller{
		placeStore: pStore,
		userStore:  uStore,
		minVersion: mVersion,
	}
}

// TODO: move identity check via interceptors on adapter layer
func (c *Controller) Authorize(ctx context.Context,
	aToken string, uuid string, ver util.Version) (*util.Achievements, error) {

	// if ver.Less(c.minVersion) {
	// 	return nil, icontroller.ErrVersionNotCompatible
	// }

	// u, err := c.userStore.GetByID(ctx, uuid)
	// if err != nil {
	// 	log.Printf("Controller.Authorize: db error: %s\n", err)
	// 	return nil, err
	// }

	// return &u.Achievements, nil
	panic("Unimplemented")
}

// TODO: move identity check via interceptors on adapter layer
func (c *Controller) GetPlaces(ctx context.Context,
	tinkUUID string, aToken string, ctg category.Category,
	gCenter util.LatLng) ([]place.Place, error) {

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
