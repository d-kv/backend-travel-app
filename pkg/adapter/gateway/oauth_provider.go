package gateway

import (
	"context"
)

type OAuthProviderI interface {
	GetUserID(ctx context.Context, accessToken string) (string, error)
}
