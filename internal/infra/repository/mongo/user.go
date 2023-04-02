// TODO: add tests
package mongo

import (
	"context"
	"errors"
	"log"

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

// GetAll returns all users.
func (u *UserStore) GetAll(ctx context.Context) ([]user.User, error) {
	cursor, err := u.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("UserStore.GetAll: db error: %s\n", err)
		return nil, err
	}

	var users []user.User
	err = cursor.All(ctx, &users) // FIXME: may be an overflow
	if err != nil {
		log.Printf("UserStore.GetAll: decoding error: %s\n", err)
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
		log.Printf("UserStore.Create: DB error: %s\n", err)
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
		log.Printf("UserStore.Delete: db error: %s\n", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Printf("UserStore.Delete: db error: %s\n", irepository.ErrUserNotFound)
		return irepository.ErrUserNotFound
	}

	if res.DeletedCount > 1 {
		log.Printf("UserStore.Delete: db error: %s\n", irepository.ErrUUIDDuplicate)
		return irepository.ErrUUIDDuplicate
	}

	return nil
}

// GetByID returns user with given UUID.
func (u *UserStore) GetByID(ctx context.Context, uuid string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"_id": uuid,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Printf("UserStore.GetByID: db error: %s\n", err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		log.Printf("UserStore.GetByID: db error: %s\n", err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		log.Printf("UserStore.GetByID: decoding error: %s\n", err)
		return nil, err
	}

	return user, nil
}

// GetByTinkoffID returns user with given Tinkoff UUID.
func (u *UserStore) GetByTinkoffID(ctx context.Context, tinkoffUUID string) (*user.User, error) {
	res := u.coll.FindOne(ctx, bson.M{
		"tinkoff_id": tinkoffUUID,
	})

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Printf("UserStore.GetByTinkoffID: db error: %s\n", err)
		return nil, irepository.ErrUserNotFound
	}

	if err != nil {
		log.Printf("UserStore.GetByTinkoffID: db error: %s\n", err)
		return nil, err
	}

	var user *user.User
	err = res.Decode(&user)
	if err != nil {
		log.Printf("UserStore.GetByTinkoffID: decoding error: %s\n", err)
		return nil, err
	}

	return user, nil
}
