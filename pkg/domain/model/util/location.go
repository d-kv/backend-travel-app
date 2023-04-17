package util

type geo struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// Location stores geospatial info.
type Location struct {
	Geo geo `bson:"geo"`
}

func (l *Location) Latitude() float64 {
	return l.Geo.Coordinates[1]
}

func (l *Location) Longitude() float64 {
	return l.Geo.Coordinates[0]
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
