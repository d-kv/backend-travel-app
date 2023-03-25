package query

import "github.com/d-kv/backend-travel-app/pkg/domain/model/util"

type Distance struct {
	SearchAreaCenter util.LatLng
	Min              int32
	Max              int32
}
