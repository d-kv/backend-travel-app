// TODO: add tests
package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/d-kv/backend-travel-app/pkg/infra/ilogger"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// UserStore with CRUD-like operations on the User object.
type UserStore struct {
	log  ilogger.LoggerI
	coll *mongo.Collection
}

var _ irepository.UserI = (*UserStore)(nil)

// NewUserStore is a default ctor.
func NewUserStore(l ilogger.LoggerI, coll *mongo.Collection) *UserStore {
	return &UserStore{
		log:  l,
		coll: coll,
	}
}

// GetAll returns all users.
func (u *UserStore) GetAll(ctx context.Context) ([]user.User, error) {
	cursor, err := u.coll.Find(ctx, bson.D{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		u.log.Info("UserStore.GetByID: %v", err)
		return nil, irepository.ErrUserNotFound
	}
	if err != nil {
		u.log.Error("UserStore.GetAll: %v", err)
		return nil, err
	}

	var users []user.User
	err = cursor.All(ctx, &users) // FIXME: may be an overflow
	if err != nil {
		u.log.Error("UserStore.GetAll: %v", err)
		return nil, err
	}

	return users, nil
}

// Create creates a new user.
//
// UUID field must be populated.
func (u *UserStore) Create(ctx context.Context, user *user.User) error {
	if user.UUID == "" {
		return irepository.ErrUUIDNotPopulated
	}

	_, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		u.log.Warn("UserStore.Create: %v", err)
		return err
	}

	return nil
}

// Delete deletes user with given UUID.
func (u *UserStore) Delete(ctx context.Context, uuid string) error {
	res, err := u.coll.DeleteOne(ctx, bson.M{
		"_id": uuid,
	})
	if err != nil {
		u.log.Warn("UserStore.Delete: %v", err)
		return err
	}

	if res.DeletedCount == 0 {
		u.log.Warn("UserStore.Delete: %v", irepository.ErrUserNotFound)
		return irepository.ErrUserNotFound
	}

	if res.DeletedCount > 1 {
		u.log.Error("UserStore.Delete: %v", irepository.ErrUUIDDuplicate)
		return irepository.ErrUUIDDuplicate
	}

	return nil
}

// Get returns user with given UUID.
func (u *UserStore) Get(ctx context.Context, uuid string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		u.log.Info("UserStore.GetByID: %v", err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		u.log.Warn("UserStore.GetByID: %v", err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		u.log.Error("UserStore.GetByID: %v", err)
		return nil, err
	}

	return user, nil
}

// GetByOAuthID returns user with given OAuth Provider ID.
func (u *UserStore) GetByOAuthID(ctx context.Context, oAuthID string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"oauth_id": oAuthID,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		u.log.Info("UserStore.GetByOAuthID: %v", err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		u.log.Warn("UserStore.GetByOAuthID: %v", err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		u.log.Error("UserStore.GetByOAuthID: %v", err)
		return nil, err
	}

	return user, nil
}

// GetByOAuthAToken returns user with given OAuth Provider Access Token.
func (u UserStore) GetByOAuthAToken(ctx context.Context, oAuthAToken string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"oauth_access_token": oAuthAToken,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		u.log.Info("UserStore.GetByOAuthAToken: %v", err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		u.log.Warn("UserStore.GetByOAuthAToken: %v", err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		u.log.Error("UserStore.GetByOAuthAToken: %v", err)
		return nil, err
	}

	return user, nil
}
