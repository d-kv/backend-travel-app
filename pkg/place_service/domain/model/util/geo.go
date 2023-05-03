package util

const (
	DefaultMinDistance = 0
	DefaultMaxDistance = 5000
)

type GeoQuery struct {
	Center *LatLng
	Min    int64
	Max    int64
}

type Options func(*GeoQuery)

func WithMin(min int64) Options {
	return func(p *GeoQuery) { p.Min = min }
}

func WithMax(max int64) Options {
	return func(p *GeoQuery) { p.Max = max }
}

// NewGeoQuery creates a new GeoQuery.
func NewGeoQuery(ll *LatLng, opts ...Options) *GeoQuery {
	g := &GeoQuery{
		Center: ll,
		Min:    DefaultMinDistance,
		Max:    DefaultMaxDistance,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}
