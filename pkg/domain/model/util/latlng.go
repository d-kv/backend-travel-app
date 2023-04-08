// TODO: use explicit naming for langitude & longitude
// For instance, we can use prefixes, such as: "lat_56.34564,lng_46.2356"
package util

import (
	"strconv"
	"strings"
)

// LatLng stores latitude & longitude.
type LatLng struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

// NewLatLng creates a new LatLng with given lat & lng values.
func NewLatLng(lat, lng float64) *LatLng {
	return &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}
}

// NewLatLngFromString creates a new LatLng from a string of the form "<latitude>,<longitude>".
func NewLatLngFromString(llStr string) (*LatLng, error) {
	ll := &LatLng{}
	if err := parseLatLngFromString(ll, llStr); err != nil {
		return nil, err
	}

	return ll, nil
}

func parseLatLngFromString(ll *LatLng, rawStr string) error {
	parts := strings.Split(rawStr, ",")

	lat, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return err
	}
	ll.Latitude = lat

	lng, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return err
	}
	ll.Longitude = lng

	return nil
}
