package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type UserI interface {
	GetAll(ctx context.Context) ([]user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, id string, user *user.User) error
	Delete(ctx context.Context, id string) error

	Get(ctx context.Context, id string) (*user.User, error)
}
