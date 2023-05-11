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

const (
	minLatitude  = -90
	maxLatitude  = 90
	minLongitude = -180
	maxLongitude = 180
)

// LatLng stores latitude & longitude.
type LatLng struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

// NewLatLng creates a new LatLng with given lat & lng values.
func NewLatLng(lat, lng float64) (*LatLng, error) {
	if !IsValidLatitude(lat) || !IsValidLongitude(lng) {
		return nil, ErrInvalidLatLng
	}
	return &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}, nil
}

func IsValidLatitude(lat float64) bool {
	return lat > minLatitude || lat < maxLatitude
}

func IsValidLongitude(lng float64) bool {
	return lng > minLongitude || lng < maxLongitude
}

// NewLatLngFromString creates a new LatLng from a string of the form "<latitude>,<longitude>".
func NewLatLngFromString(llStr string) (*LatLng, error) {
	lat, lng, err := parseLatLngFromString(llStr)
	if err != nil {
		return nil, ErrUnableToParseLatLng
	}

	if !IsValidLatitude(lat) || !IsValidLongitude(lng) {
		return nil, ErrInvalidLatLng
	}

	return &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}, nil
}

// NewLatLngFromRString creates a new LatLng from a string of the form "<longitude>,<latitude>".
func NewLatLngFromRString(llStr string) (*LatLng, error) {
	lng, lat, err := parseLatLngFromString(llStr)
	if err != nil {
		return nil, ErrUnableToParseLatLng
	}

	if !IsValidLatitude(lat) || !IsValidLongitude(lng) {
		return nil, ErrInvalidLatLng
	}

	return &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}, nil
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
