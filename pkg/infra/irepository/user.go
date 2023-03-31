package irepository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type UserI interface {
	GetAll(context.Context) ([]*user.User, error)
	Create(context.Context, *user.User) error
	Delete(context.Context, string) error

	GetByID(context.Context, string) (*user.User, error)
}
