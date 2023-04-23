// TODO: add tests
package mongo

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

// UserStore with CRUD-like operations on the User object.
type UserStore struct {
	coll *mongo.Collection
}

var _ irepository.UserI = (*UserStore)(nil)

// NewUserStore is a default ctor.
func NewUserStore(coll *mongo.Collection) *UserStore {
	return &UserStore{
		coll: coll,
	}
}

// Users returns all users.
func (u *UserStore) Users(ctx context.Context) ([]user.User, error) {
	cursor, err := u.coll.Find(ctx, bson.D{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrUserNotFound
	}
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	var users []user.User
	err = cursor.All(ctx, &users) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Err(err)
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
		log.Warn().
			Err(err)
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
		log.Warn().
			Err(err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Warn().
			Err(err)
		return irepository.ErrUserNotFound
	}

	if res.DeletedCount > 1 {
		log.Error().
			Err(err)
		return irepository.ErrUUIDDuplicate
	}

	return nil
}

func (u *UserStore) Update(ctx context.Context, uuid string, user *user.User) error {
	_, err := u.coll.ReplaceOne(ctx, bson.M{"_id": uuid}, user)
	if err != nil {
		log.Warn().
			Err(err)
		return err
	}
	return nil
}

// User returns user with given UUID.
func (u *UserStore) User(ctx context.Context, uuid string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Err(err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		log.Warn().
			Err(err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return user, nil
}
