package util

type geo struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Location struct {
	Geo geo `bson:"geo"`
}
