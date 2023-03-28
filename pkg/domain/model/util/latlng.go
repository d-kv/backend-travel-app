package util

type LatLng struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

func NewLatLng(lat, lng float64) LatLng {
	return LatLng{
		Latitude:  lat,
		Longitude: lng,
	}
}
