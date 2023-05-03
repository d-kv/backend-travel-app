package category

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

// MainCategory is a enum type.
type MainCategory int32

type bsonMainCategory struct {
	String string
}

var _ bson.Marshaler = (*MainCategory)(nil)
var _ bson.Unmarshaler = (*MainCategory)(nil)
var _ json.Marshaler = (*MainCategory)(nil)
var _ json.Unmarshaler = (*MainCategory)(nil)

func (m MainCategory) String() string {
	return mainCategoryName[m]
}

func MainCategoryFromString(mcRaw string) MainCategory {
	return mainCategoryValue[mcRaw]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	MC_UNSPECIFIED MainCategory = iota
	MC_CULTURE
	MC_ENTERTAINMENT
	MC_FOOD
	MC_HOSPITALITY
)

var (
	mainCategoryName = map[MainCategory]string{ //nolint:gochecknoglobals // Using global maps for enums
		MC_UNSPECIFIED:   "MC_UNSPECIFIED",
		MC_CULTURE:       "MC_CULTURE",
		MC_ENTERTAINMENT: "MC_ENTERTAINMENT",
		MC_FOOD:          "MC_FOOD",
		MC_HOSPITALITY:   "MC_HOSPITALITY",
	}
	mainCategoryValue = map[string]MainCategory{ //nolint:gochecknoglobals // Using global maps for enums
		"MC_UNSPECIFIED":   MC_UNSPECIFIED,
		"MC_CULTURE":       MC_CULTURE,
		"MC_ENTERTAINMENT": MC_ENTERTAINMENT,
		"MC_FOOD":          MC_FOOD,
		"MC_HOSPITALITY":   MC_HOSPITALITY,
	}
)

func (m MainCategory) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bsonMainCategory{
		String: mainCategoryName[m],
	})
}

func (m *MainCategory) UnmarshalBSON(data []byte) error {
	var bsonRepr bsonMainCategory
	err := bson.Unmarshal(data, &bsonRepr)
	if err != nil {
		return err
	}
	*m = mainCategoryValue[bsonRepr.String]
	return nil
}

func (m MainCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(mainCategoryName[m])
}

func (m *MainCategory) UnmarshalJSON(data []byte) error {
	var jsonRepr string

	if err := json.Unmarshal(data, &jsonRepr); err != nil {
		return err
	}
	*m = mainCategoryValue[jsonRepr]

	return nil
}
