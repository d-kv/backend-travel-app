package icontroller

import (
	"context"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

type ControllerI interface {
	GetAchievements(context.Context, string) (*util.Achievements, error)
	GetPlaces(ctx context.Context) ([]place.Place, error)
}
