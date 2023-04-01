package place

import (
	"time"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type Place struct {
	// TODO: add bill, opening_hours & rating
	UUID        string `bson:"_id"`
	Address     string `bson:"address"`
	Name        string `bson:"name"`
	Description string `bson:"description,omitempty"`
	Phone       string `bson:"phone,omitempty"`
	URL         string `bson:"url, omitempty"`

	LatLng   util.LatLng       `bson:"inline"`
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
		p.LatLng.Latitude = lat
		p.LatLng.Longitude = lng
	}
}

func WithCategory(cat category.Category) Options {
	return func(p *Place) { p.Category = cat }
}

func WithLifetime(lt time.Duration) Options {
	return func(p *Place) { p.Lifetime = lt }
}

func WithRecord(cAt, uAt time.Time) Options {
	return func(p *Place) {
		p.CreatedAt = cAt
		p.UpdatedAt = uAt
	}
}

func New(opts ...Options) *Place {
	p := &Place{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
