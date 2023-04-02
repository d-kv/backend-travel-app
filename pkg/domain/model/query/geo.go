package query

import "github.com/d-kv/backend-travel-app/pkg/domain/model/util"

type Geo struct {
	Center *util.LatLng
	Min    int64
	Max    int64
}
