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

type builderI interface {
	UUID(val string) builderI
	Address(val string) builderI
	Name(val string) builderI
	Description(val string) builderI
	Phone(val string) builderI

	LatLng(lat, lng float64) builderI
	Category(cat category.Category) builderI

	Lifetime(val time.Duration) builderI
	Record(cAt, uAt time.Time) builderI
}

type builder struct {
	uuid        string
	address     string
	name        string
	description string
	phone       string

	latLng   util.LatLng
	category category.Category

	lifetime time.Duration
	record   util.Record
}

var _ builderI = builder{}

func NewBuilder() builderI {
	return builder{}
}

func (b builder) UUID(val string) builderI {
	b.uuid = val
	return b
}

func (b builder) Address(val string) builderI {
	b.address = val
	return b
}

func (b builder) Name(val string) builderI {
	b.name = val
	return b
}

func (b builder) Description(val string) builderI {
	b.description = val
	return b
}

func (b builder) Phone(val string) builderI {
	b.phone = val
	return b
}

func (b builder) LatLng(lat float64, lng float64) builderI {
	b.latLng = util.LatLng{
		Latitude:  lat,
		Longitude: lng,
	}
	return b
}

func (b builder) Category(val category.Category) builderI {
	b.category = val
	return b
}

func (b builder) Lifetime(val time.Duration) builderI {
	b.lifetime = val
	return b
}

func (b builder) Record(cAt, uAt time.Time) builderI {
	b.record = util.Record{
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}
	return b
}

func (b builder) Build() Place {
	return Place{
		UUID:        b.uuid,
		Address:     b.address,
		Name:        b.name,
		Description: b.description,
		Phone:       b.phone,
		LatLng:      b.latLng,
		Category:    b.category,
		Lifetime:    b.lifetime,
		Record:      b.record,
	}
}
