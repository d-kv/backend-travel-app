// TODO: add tests
package controllerv0

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=PlaceProvider --output=mock --case=underscore --disable-version-string --outpkg=mock
//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=UserProvider --output=mock --case=underscore --disable-version-string --outpkg=mock
//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=OAuthProvider --output=mock --case=underscore --disable-version-string --outpkg=mock
//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=TokenCache --output=mock --case=underscore --disable-version-string --outpkg=mock

type (
	PlaceProvider interface {
		PlacesByDistance(ctx context.Context, geoQ *query.Geo, skipN int64, resN int64) ([]place.Place, error)
	}
	UserProvider interface {
		Update(ctx context.Context, uuid string, user *user.User) error
		User(ctx context.Context, uuid string) (*user.User, error)
	}

	OAuthProvider interface {
		UserID(ctx context.Context, accessToken string) (userUUID string, err error)
	}

	TokenCache interface {
		SetUserID(ctx context.Context, refreshToken, userUUID string) error
		UserID(ctx context.Context, refreshToken string) (userUUID string, err error)
	}
)

type Controller struct {
	placeProvider PlaceProvider
	userProvider  UserProvider
	tokenCache    TokenCache
	oAuthProvider OAuthProvider
}

// New is a default ctor for Controller.
func New(placeP PlaceProvider, userP UserProvider, tokenC TokenCache, oAuthP OAuthProvider) *Controller {
	return &Controller{
		placeProvider: placeP,
		userProvider:  userP,
		oAuthProvider: oAuthP,
		tokenCache:    tokenC,
	}
}
