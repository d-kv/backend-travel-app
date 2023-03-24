package place

import (
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"

	"time"
)

type Place struct {
	// TODO: add bill, opening_hours & rating
	UUID        string
	Address     string
	Name        string
	Description string
	Phone       string

	LatLng   util.LatLng
	Category category.Category

	Lifetime time.Duration
	Record   util.Record
}

type PlaceOpts func(*Place)

func WithUUID(uuid string) PlaceOpts {
	return func(p *Place) { p.UUID = uuid }
}

func WithAddress(addr string) PlaceOpts {
	return func(p *Place) { p.Address = addr }
}

func WithName(name string) PlaceOpts {
	return func(p *Place) { p.Name = name }
}

func WithDescription(desc string) PlaceOpts {
	return func(p *Place) { p.Description = desc }
}

func WithPhone(phone string) PlaceOpts {
	return func(p *Place) { p.Phone = phone }
}

func WithLatLng(lat float64, lng float64) PlaceOpts {
	return func(p *Place) {
		p.LatLng.Latitude = lat
		p.LatLng.Longitude = lng
	}
}

func WithCategory(cat category.Category) PlaceOpts {
	return func(p *Place) { p.Category = cat }
}

func WithLifetime(lt time.Duration) PlaceOpts {
	return func(p *Place) { p.Lifetime = lt }
}

func WithRecord(cAt, uAt time.Time) PlaceOpts {
	return func(p *Place) {
		p.Record.CreatedAt = cAt
		p.Record.UpdatedAt = uAt
	}
}

func NewPlace(opts ...PlaceOpts) *Place {
	p := &Place{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
