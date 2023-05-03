package repository

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/user-service/model"
)

type UserProvider interface {
	Users(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error

	User(ctx context.Context, id string) (*model.User, error)
}
