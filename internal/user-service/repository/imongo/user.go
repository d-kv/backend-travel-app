package imongo

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/d-kv/backend-travel-app/pkg/user-service/model"
	"github.com/d-kv/backend-travel-app/pkg/user-service/repository"
)

// UserStore with CRUD-like operations on the User object.
type UserStore struct {
	coll *mongo.Collection
}

var _ repository.UserProvider = (*UserStore)(nil)

// New is a default ctor.
func New(coll *mongo.Collection) *UserStore {
	return &UserStore{
		coll: coll,
	}
}

// Users returns all users.
func (u *UserStore) Users(ctx context.Context) ([]model.User, error) {
	const mName = "UserStore.Users"

	cursor, err := u.coll.Find(ctx, bson.D{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Info().
			Str("method", mName).
			Err(err).
			Msg("user db is empty")

		return nil, repository.ErrUserNotFound
	}
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return nil, err
	}

	var users []model.User
	err = cursor.All(ctx, &users) // FIXME: may be an overflow
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error while decoding")

		return nil, err
	}

	return users, nil
}

// Create creates a new user.
//
// UUID field must be populated.
func (u *UserStore) Create(ctx context.Context, user *model.User) error {
	const mName = "UserStore.Create"

	if user.UUID == "" {
		return repository.ErrUUIDNotPopulated
	}

	_, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return err
	}

	return nil
}

// Delete deletes user with given UUID.
func (u *UserStore) Delete(ctx context.Context, uuid string) error {
	const mName = "UserStore.Delete"

	res, err := u.coll.DeleteOne(ctx, bson.M{
		"_id": uuid,
	})
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return err
	}

	if res.DeletedCount == 0 {
		return repository.ErrUserNotFound
	}

	return nil
}

func (u *UserStore) Update(ctx context.Context, uuid string, user *model.User) error {
	const mName = "UserStore.Update"

	_, err := u.coll.ReplaceOne(ctx, bson.M{"_id": uuid}, user)
	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return err
	}
	return nil
}

// User returns user with given UUID.
func (u *UserStore) User(ctx context.Context, uuid string) (*model.User, error) {
	const mName = "UserStore.User"

	res := u.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository.ErrUserNotFound
	}

	if err != nil {
		log.Warn().
			Str("method", mName).
			Err(err).
			Msg("error from mongoDB driver")

		return nil, err
	}

	var user *model.User
	err = res.Decode(&user)
	if err != nil {
		log.Error().
			Str("method", mName).
			Err(err).
			Msg("error while decoding")

		return nil, err
	}

	return user, nil
}
