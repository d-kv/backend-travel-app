// TODO: add tests
package controllerv0

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/adapter/igateway"
	icontrollerv0 "github.com/d-kv/backend-travel-app/pkg/app/icontroller/v0"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// Controller defines a place service controller.
type Controller struct {
	placeStore    irepository.PlaceI
	userStore     irepository.UserI
	oAuthProvider igateway.OAuthProviderI
}

var _ icontrollerv0.ControllerI = (*Controller)(nil)

// New is a default ctor for Controller.
func New(pStore irepository.PlaceI, uStore irepository.UserI) *Controller {
	return &Controller{
		placeStore: pStore,
		userStore:  uStore,
	}
}

func (c *Controller) GetUser(ctx context.Context, oAuthAToken string) (*user.User, error) {
	u, err := c.userStore.GetByOAuthAToken(ctx, oAuthAToken)
	if errors.Is(err, irepository.ErrUserNotFound) {
		oAuthID, err := c.oAuthProvider.GetUserID(ctx, oAuthAToken)
		if err != nil {
			log.Info().Msgf("Controller.GetUser: %v", err)
			return nil, err
		}

		newU := user.New(
			user.WithOAuthAToken(oAuthAToken),
			user.WithOAuthID(oAuthID),
		)

		err = c.userStore.Create(ctx, newU)
		if err != nil {
			log.Info().Msgf("Controller.GetUser: %v", err)
		}
		return newU, nil
	} else if err != nil {
		log.Info().Msgf("Controller.GetUser: %v", err)
		return nil, err
	}

	return u, nil
}

func (c *Controller) GetPlaces(ctx context.Context, gCenter *util.LatLng) ([]place.Place, error) {
	geoQ := query.Geo{ // TODO: receive min & max parameters from request
		Center: gCenter,
	}

	places, err := c.placeStore.GetNearby(ctx, &geoQ)
	if err != nil {
		log.Info().Msgf("Controller.GetPlaces: %v\n", err)
		return nil, err
	}

	return places, nil
}
