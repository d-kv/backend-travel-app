package user

import (
	"time"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type User struct {
	UUID               string `bson:"_id"`
	TinkoffID          string `bson:"tinkoff_id"`
	TinkoffAccessToken string `bson:"tinkoff_access_token"`

	Premium      bool              `bson:"premium"`
	Tester       bool              `bson:"tester"`
	Admin        bool              `bson:"admin"`
	Blocked      bool              `bson:"blocked"`
	Achievements util.Achievements `bson:"achievements"`

	LastActivity time.Time `bson:"last_activity"`
}

type Options func(*User)

func WithUUID(uuid string) Options {
	return func(u *User) { u.UUID = uuid }
}

func WithTinkoffID(tinkoffID string) Options {
	return func(u *User) { u.TinkoffID = tinkoffID }
}

func WithTinkoffAccessToken(tinkoffAccessToken string) Options {
	return func(u *User) { u.TinkoffAccessToken = tinkoffAccessToken }
}

func WithPremium(premium bool) Options {
	return func(u *User) { u.Premium = premium }
}

func WithTester(tester bool) Options {
	return func(u *User) { u.Tester = tester }
}

func WithAdmin(admin bool) Options {
	return func(u *User) { u.Admin = admin }
}

func WithBlocked(blocked bool) Options {
	return func(u *User) { u.Blocked = blocked }
}

func WithLastActivity(lastActivity time.Time) Options {
	return func(u *User) { u.LastActivity = lastActivity }
}

func New(opts ...Options) *User {
	u := &User{}

	for _, opt := range opts {
		opt(u)
	}

	return u
}
