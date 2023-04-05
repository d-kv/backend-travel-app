package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type UserI interface {
	GetAll(ctx context.Context) ([]user.User, error)
	Create(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*user.User, error)
	GetByOAuthID(ctx context.Context, id string) (*user.User, error)
	GetByOAuthAToken(ctx context.Context, accessToken string) (*user.User, error)
}
