package place

import (
	"time"

	"github.com/google/uuid"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type Place struct {
	// TODO: add bill, opening_hours & rating
	// TODO: move address to Location field & parse it there into City, Street, etc
	UUID        string `bson:"_id"`
	Address     string `bson:"address"`
	Name        string `bson:"name"`
	Description string `bson:"description,omitempty"`
	Phone       string `bson:"phone,omitempty"`
	URL         string `bson:"url,omitempty"`

	Location util.Location     `bson:"location"`
	Category category.Category `bson:"category"`

	Lifetime  time.Duration `bson:"lifetime"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type Options func(*Place)

func WithUUID(uuid string) Options {
	return func(p *Place) { p.UUID = uuid }
}

func WithAddress(addr string) Options {
	return func(p *Place) { p.Address = addr }
}

func WithName(name string) Options {
	return func(p *Place) { p.Name = name }
}

func WithDescription(desc string) Options {
	return func(p *Place) { p.Description = desc }
}

func WithPhone(phone string) Options {
	return func(p *Place) { p.Phone = phone }
}

func WithURL(url string) Options {
	return func(p *Place) { p.URL = url }
}

func WithLatLng(lat float64, lng float64) Options {
	return func(p *Place) {
		p.Location.Geo.Type = "Point"
		p.Location.Geo.Coordinates = []float64{lng, lat}
	}
}

func WithCategory(cat category.Category) Options {
	return func(p *Place) { p.Category = cat }
}

func WithLifetime(ttl time.Duration) Options {
	return func(p *Place) { p.Lifetime = ttl }
}

// New creates a new Place.
func New(opts ...Options) *Place {
	p := &Place{
		UUID:        uuid.New().String(),
		Address:     "",
		Name:        "",
		Description: "",
		Phone:       "",
		URL:         "",
		Location:    *util.NewLocation(*util.NewLatLng(0, 0)),
		Category:    category.Category{},
		Lifetime:    0,
		// FIXME: make compatible with MongoDB precision
		// CreatedAt:   time.Now().UTC(),
		// UpdatedAt:   time.Now().UTC(),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
