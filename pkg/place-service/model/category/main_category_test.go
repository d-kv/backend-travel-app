package category_test

import (
	"encoding/json"
	"testing"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMainCategorySerialization(t *testing.T) {
	assert := assert.New(t)

	sFmts := []struct {
		Name      string
		Marshal   func(v any) ([]byte, error)
		Unmarshal func(d []byte, v any) error
	}{
		{"bson", bson.Marshal, bson.Unmarshal},
		{"json", json.Marshal, json.Unmarshal},
	}

	inCtgs := []category.Main{
		category.MC_UNSPECIFIED,
		category.MC_CULTURE,
		category.MC_ENTERTAINMENT,
		category.MC_FOOD,
		category.MC_HOSPITALITY,
	}
	for _, sFmt := range sFmts {
		for _, c := range inCtgs {
			jsonRepr, err := sFmt.Marshal(c)
			assert.NoErrorf(err, "must be serialized into %s without errors", sFmt.Name)

			var gotCtg category.Main
			err = sFmt.Unmarshal(jsonRepr, &gotCtg)
			assert.NoErrorf(err, "must be deserialized from %s without errors", sFmt.Name)

			assert.Equal(c, gotCtg, "must be the same")
		}
	}
}
