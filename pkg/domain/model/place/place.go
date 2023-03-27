package place

import (
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"

	"github.com/google/uuid"

	"time"
)

type Place struct {
	// TODO: add bill, opening_hours & rating
	UUID        uuid.UUID
	Address     string
	Name        string
	Description string
	Phone       string

	LatLng   util.LatLng
	Category category.Category

	Lifetime time.Duration
	Record   util.Record
}

type Options func(*Place)

func WithUUID(uuid uuid.UUID) Options {
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
		p.Record.CreatedAt = cAt
		p.Record.UpdatedAt = uAt
	}
}

func NewPlace(opts ...Options) *Place {
	p := &Place{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
