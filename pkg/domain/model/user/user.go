package user

import (
	"time"

	"github.com/google/uuid"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type User struct {
	// TODO: split Account repo & User repo
	UUID        string `bson:"_id"`
	OAuthID     string `bson:"oauth_id"`
	OAuthAToken string `bson:"oauth_access_token"`

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

func WithOAuthID(oAuthID string) Options {
	return func(u *User) { u.OAuthID = oAuthID }
}

func WithOAuthAToken(oAuthAToken string) Options {
	return func(u *User) { u.OAuthAToken = oAuthAToken }
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
	u := &User{
		UUID:         uuid.New().String(),
		OAuthID:      "",
		OAuthAToken:  "",
		Premium:      false,
		Tester:       false,
		Admin:        false,
		Blocked:      false,
		Achievements: *util.NewAchievements(""),
		// FIXME: make compatible with MongoDB precision
		// LastActivity:           time.Time{},
	}

	for _, opt := range opts {
		opt(u)
	}

	return u
}
