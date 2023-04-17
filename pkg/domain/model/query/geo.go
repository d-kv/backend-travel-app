package query

import (
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

const (
	DefaultMinDistance = 0
	DefaultMaxDistance = 5000
)

type Geo struct {
	Center *util.LatLng
	Min    int64
	Max    int64
}

type Options func(*Geo)

func WithMin(min int64) Options {
	return func(p *Geo) { p.Min = min }
}

func WithMax(max int64) Options {
	return func(p *Geo) { p.Max = max }
}

// New creates a new Geo.
func New(ll *util.LatLng, opts ...Options) *Geo {
	g := &Geo{
		Center: ll,
		Min:    DefaultMinDistance,
		Max:    DefaultMaxDistance,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}
