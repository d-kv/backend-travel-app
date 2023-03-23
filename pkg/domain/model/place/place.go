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

	lifetime time.Duration
	Record   record
}

type record struct {
	Lifetime  time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
}
