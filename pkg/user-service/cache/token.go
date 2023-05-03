package cache

import "context"

type TokenCacher interface {
	SetUserID(ctx context.Context, rToken, userUUID string) error
	UserID(ctx context.Context, rToken string) (string, error)
}
