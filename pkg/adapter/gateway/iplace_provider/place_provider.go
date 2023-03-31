package iplaceprovider

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
)

type PlaceProviderI interface {
	Get(context.Context, string) ([]*place.Place, error)
}
