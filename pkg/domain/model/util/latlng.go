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

// ParseLatLngFromString populates LatLng from a string of the form "<latitude>,<longitude>".
func ParseLatLngFromString(ll *LatLng, rawStr string) error {
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
