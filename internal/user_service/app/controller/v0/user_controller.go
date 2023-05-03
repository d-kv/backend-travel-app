// TODO: add tests
package iuser_ctrl_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/user_service/domain/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=UserProvider --output=mock --case=underscore --disable-version-string --outpkg=mock
//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=OAuthProvider --output=mock --case=underscore --disable-version-string --outpkg=mock
//go:generate go run github.com/vektra/mockery/v2@v2.25.1 --name=TokenCache --output=mock --case=underscore --disable-version-string --outpkg=mock

type (
	UserProvider interface {
		Update(ctx context.Context, uuid string, user *model.User) error
		User(ctx context.Context, uuid string) (*model.User, error)
	}

	OAuthProvider interface {
		UserID(ctx context.Context, accessToken string) (userUUID string, err error)
	}

	TokenCache interface {
		SetUserID(ctx context.Context, refreshToken, userUUID string) error
		UserID(ctx context.Context, refreshToken string) (userUUID string, err error)
	}
)

type UserController struct {
	userProvider  UserProvider
	tokenCache    TokenCache
	oAuthProvider OAuthProvider
}

// New is a default ctor for UserController.
func New(userP UserProvider, tokenC TokenCache, oAuthP OAuthProvider) *UserController {
	return &UserController{
		userProvider:  userP,
		oAuthProvider: oAuthP,
		tokenCache:    tokenC,
	}
}
