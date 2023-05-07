package util

const (
	DefaultMinDistance = 0
	DefaultMaxDistance = 5000
)

type GeoToken struct {
	Center *LatLng
	Min    int64
	Max    int64
}

type Options func(*GeoToken)

func WithMin(min int64) Options {
	return func(p *GeoToken) { p.Min = min }
}

func WithMax(max int64) Options {
	return func(p *GeoToken) { p.Max = max }
}

// NewGeoToken creates a new GeoQuery.
func NewGeoToken(ll *LatLng, opts ...Options) *GeoToken {
	g := &GeoToken{
		Center: ll,
		Min:    DefaultMinDistance,
		Max:    DefaultMaxDistance,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}
