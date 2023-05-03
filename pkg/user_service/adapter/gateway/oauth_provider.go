package gateway

import (
	"context"
)

type OAuthProviderI interface {
	UserID(ctx context.Context, accessToken string) (string, error)
}
