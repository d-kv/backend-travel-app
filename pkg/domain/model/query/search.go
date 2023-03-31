package query

import (
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
)

type Search struct {
	Categories []category.Classification
	Distance   Distance
}
