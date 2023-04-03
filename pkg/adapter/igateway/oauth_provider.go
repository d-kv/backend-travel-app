package igateway

import (
	"context"
)

type OAuthProviderI interface {
	GetUserID(context.Context, string) (string, error)
}
