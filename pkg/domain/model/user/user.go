package user

import (
	"time"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
	"github.com/google/uuid"
)

type User struct {
	UUID                   string `bson:"_id"`
	IdentityProviderID     string `bson:"identity_provider_id"`
	IdentityProviderAToken string `bson:"identity_provider_access_token"`

	Premium      bool              `bson:"premium"`
	Tester       bool              `bson:"tester"`
	Admin        bool              `bson:"admin"`
	Blocked      bool              `bson:"blocked"`
	Achievements util.Achievements `bson:"achievements"`

	LastActivity time.Time `bson:"last_activity"`
}

func NewDefault(identityProviderID, identityProviderAToken string) *User {
	return &User{
		UUID:                   uuid.New().String(),
		IdentityProviderID:     identityProviderID,
		IdentityProviderAToken: identityProviderAToken,
		Premium:                false,
		Tester:                 false,
		Admin:                  false,
		Blocked:                false,
		Achievements:           util.Achievements{},
		LastActivity:           time.Now(),
	}
}

type Options func(*User)

func WithUUID(uuid string) Options {
	return func(u *User) { u.UUID = uuid }
}

func WithIdentityProviderID(identityProviderID string) Options {
	return func(u *User) { u.IdentityProviderID = identityProviderID }
}

func WithIdentityProviderAToken(identityProviderAToken string) Options {
	return func(u *User) { u.IdentityProviderAToken = identityProviderAToken }
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
		UUID:                   uuid.New().String(),
		IdentityProviderID:     "",
		IdentityProviderAToken: "",
		Premium:                false,
		Tester:                 false,
		Admin:                  false,
		Blocked:                false,
		Achievements:           *util.NewAchievements(""),
		// FIXME: make compatible with MongoDB precision
		// LastActivity:           time.Time{},
	}

	for _, opt := range opts {
		opt(u)
	}

	return u
}
