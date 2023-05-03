package category_test

import (
	"encoding/json"
	"testing"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCategorySerialization(t *testing.T) {
	assert := assert.New(t)

	sFmts := []struct {
		Name      string
		Marshal   func(v any) ([]byte, error)
		Unmarshal func(d []byte, v any) error
	}{
		{"bson", bson.Marshal, bson.Unmarshal},
		{"json", json.Marshal, json.Unmarshal},
	}

	inCtgs := []*category.Category{
		category.New(category.WithMainCategories( // With 2 levels of categories
			category.MC_CULTURE,
			category.MC_ENTERTAINMENT,
		), category.WithSubCategories(
			category.SC_APARTMENTS,
			category.SC_AMUSEMENT_PARK,
			category.SC_BAR,
		)),

		category.New(category.WithMainCategories( // With main category only
			category.MC_CULTURE,
		)),

		category.New(), // Without any specified category
	}
	for _, sFmt := range sFmts {
		for _, c := range inCtgs {
			jsonRepr, err := sFmt.Marshal(c)
			assert.NoErrorf(err, "must be serialized into %s without errors", sFmt.Name)

			var gotCtg *category.Category
			err = sFmt.Unmarshal(jsonRepr, &gotCtg)
			assert.NoErrorf(err, "must be deserialized from %s without errors", sFmt.Name)

			assert.Equal(c, gotCtg, "must be the same")
		}
	}
}

func BenchmarkRichCategoryCreation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_ = category.New(
			category.WithMainCategories(
				category.MC_UNSPECIFIED,
				category.MC_CULTURE,
				category.MC_ENTERTAINMENT,
				category.MC_FOOD,
				category.MC_HOSPITALITY),
			category.WithSubCategories(
				category.SC_UNSPECIFIED,
				category.SC_RUSSIAN_CUISINE,
				category.SC_ITALIAN_CUISINE,
				category.SC_APARTMENTS,
				category.SC_BOWLING,
				category.SC_CAMPING,
				category.SC_GALLERY,
				category.SC_AMUSEMENT_PARK,
				category.SC_ARCHITECTURAL_MONUMENTS,
				category.SC_BEER_HOUSE,
				category.SC_PAB,
				category.SC_VEGAN_MENU,
				category.SC_OPEN_MIC,
				category.SC_NIGHTCLUB,
				category.SC_COFFEE_HOUSE,
				category.SC_LIBRARY,
				category.SC_RESORT,
				category.SC_MOTEL,
				category.SC_CONFECTIONERY,
				category.SC_JAPANESE_CUISINE,
				category.SC_TRAMPOLINE_PARK,
				category.SC_THEATRE,
				category.SC_WATER_PARK,
				category.SC_QUEST_ROOM,
				category.SC_FESTIVAL,
				category.SC_KAFE,
				category.SC_MUSEUM,
				category.SC_GEORGIAN_CUISINE,
				category.SC_HOTEL,
				category.SC_BILLIARD_CLUB,
				category.SC_CINEMA,
				category.SC_AMERICAN_CUISINE,
				category.SC_BAR,
				category.SC_STEAK,
				category.SC_HOSTEL),
		)
	}
}

func BenchmarkEmptyCategoryCreation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_ = category.New()
	}
}
