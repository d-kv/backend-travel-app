package userstore

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type UserProvider interface {
	Users(ctx context.Context) ([]user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, id string, user *user.User) error
	Delete(ctx context.Context, id string) error

	User(ctx context.Context, id string) (*user.User, error)
}
