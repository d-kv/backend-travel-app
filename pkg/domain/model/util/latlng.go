// TODO: use explicit naming for langitude & longitude
// For instance, we can use prefixes, such as: "lat_56.34564,lng_46.2356"
package util

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrUnableToParseLatLng = errors.New("unable to parse LatLng from string")
	ErrInvalidLatLng       = errors.New("latLng parameter is invalid")
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
	lat, lng, err := parseLatLngFromString(llStr)
	if err != nil {
		return nil, ErrUnableToParseLatLng
	}

	if lat > 90 || lat < -90 {
		return nil, ErrInvalidLatLng
	}

	if lng > 180 || lng < -180 {
		return nil, ErrInvalidLatLng
	}

	ll := &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}

	return ll, nil
}

// NewLatLngFromRString creates a new LatLng from a string of the form "<longitude>,<latitude>".
func NewLatLngFromRString(llStr string) (*LatLng, error) {
	lng, lat, err := parseLatLngFromString(llStr)
	if err != nil {
		return nil, ErrUnableToParseLatLng
	}

	if lat > 90 || lat < -90 {
		return nil, ErrInvalidLatLng
	}

	if lng > 180 || lng < -180 {
		return nil, ErrInvalidLatLng
	}

	ll := &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}

	return ll, nil
}

func parseLatLngFromString(rawStr string) (float64, float64, error) {
	parts := strings.Split(rawStr, ",")
	if len(parts) != 2 { //nolint:gomnd // latitude & longitude
		return 0, 0, ErrUnableToParseLatLng
	}

	l, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), bitSize)
	if err != nil {
		return 0, 0, err
	}

	ll, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), bitSize)
	if err != nil {
		return 0, 0, err
	}

	return l, ll, nil
}
