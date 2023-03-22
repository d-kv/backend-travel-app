package place

import (
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"

	"time"
)

type Place struct {
	// TODO: add bill, opening_hours & rating
	UUID        string
	Address     string
	Name        string
	Description string
	Phone       string
	Latitude    float64
	Longitude   float64
	Category    category.Category

	Lifetime  time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
}
