package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// TODO: split Account repo & User repo
	UUID string `bson:"_id"`

	Premium      bool          `bson:"premium"`
	Tester       bool          `bson:"tester"`
	Admin        bool          `bson:"admin"`
	Blocked      bool          `bson:"blocked"`
	Achievements []Achievement `bson:"achievements"`

	LastActivity time.Time `bson:"last_activity"`
}

type Options func(*User)

func WithUUID(uuid string) Options {
	return func(u *User) { u.UUID = uuid }
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

func WithAchievements(achs []Achievement) Options {
	return func(u *User) { u.Achievements = achs }
}

func WithLastActivity(lastActivity time.Time) Options {
	return func(u *User) { u.LastActivity = lastActivity }
}

func New(opts ...Options) *User {
	u := &User{
		UUID:    uuid.New().String(),
		Premium: false,
		Tester:  false,
		Admin:   false,
		Blocked: false,
		// FIXME: make compatible with MongoDB precision
		// LastActivity:           time.Time{},
	}

	for _, opt := range opts {
		opt(u)
	}

	return u
}
