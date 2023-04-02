package util

import (
	"strconv"
	"strings"
)

type LatLng struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

func NewLatLng(lat, lng float64) *LatLng {
	return &LatLng{
		Latitude:  lat,
		Longitude: lng,
	}
}

func NewLatLngFromString(llStr string) (*LatLng, error) {
	ll := &LatLng{}
	if err := parseLatLngFromString(ll, llStr); err != nil {
		return nil, err
	}

	return ll, nil
}

// parseLatLngFromString populates LatLng from a string of the form "<latitude>,<longitude>".
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
