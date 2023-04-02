package util

type geo struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Location struct {
	Geo geo `bson:"geo"`
}

func NewLocation(ll LatLng) *Location {
	return &Location{
		Geo: geo{
			Type:        "Point",
			Coordinates: []float64{ll.Longitude, ll.Latitude},
		},
	}
}
