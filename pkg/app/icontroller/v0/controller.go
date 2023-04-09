package icontrollerv0

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type ControllerI interface {
	GetUser(ctx context.Context, accessToken string) (*user.User, error)
	GetPlaces(ctx context.Context, ll *util.LatLng) ([]place.Place, error)
}