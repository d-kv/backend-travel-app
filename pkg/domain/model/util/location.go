package util

type geo struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// Location stores geospatial info.
type Location struct {
	Geo geo `bson:"geo"`
}

// NewLocation creates a new Location.
func NewLocation(ll LatLng) *Location {
	return &Location{
		Geo: geo{
			Type:        "Point",
			Coordinates: []float64{ll.Longitude, ll.Latitude},
		},
	}
}
