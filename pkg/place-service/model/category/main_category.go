package category

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

// Main is a enum type.
type Main int32

type bsonMainCategory struct {
	String string
}

var _ bson.Marshaler = (*Main)(nil)
var _ bson.Unmarshaler = (*Main)(nil)
var _ json.Marshaler = (*Main)(nil)
var _ json.Unmarshaler = (*Main)(nil)

func (m Main) String() string {
	return mainCategoryName[m]
}

func MainCategoryFromString(mcRaw string) Main {
	return mainCategoryValue[mcRaw]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	MC_UNSPECIFIED Main = iota
	MC_CULTURE
	MC_ENTERTAINMENT
	MC_FOOD
	MC_HOSPITALITY
)

var (
	mainCategoryName = map[Main]string{ //nolint:gochecknoglobals // Using global maps for enums
		MC_UNSPECIFIED:   "MC_UNSPECIFIED",
		MC_CULTURE:       "MC_CULTURE",
		MC_ENTERTAINMENT: "MC_ENTERTAINMENT",
		MC_FOOD:          "MC_FOOD",
		MC_HOSPITALITY:   "MC_HOSPITALITY",
	}
	mainCategoryValue = map[string]Main{ //nolint:gochecknoglobals // Using global maps for enums
		"MC_UNSPECIFIED":   MC_UNSPECIFIED,
		"MC_CULTURE":       MC_CULTURE,
		"MC_ENTERTAINMENT": MC_ENTERTAINMENT,
		"MC_FOOD":          MC_FOOD,
		"MC_HOSPITALITY":   MC_HOSPITALITY,
	}
)

func (m Main) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bsonMainCategory{
		String: mainCategoryName[m],
	})
}

func (m *Main) UnmarshalBSON(data []byte) error {
	var bsonRepr bsonMainCategory
	err := bson.Unmarshal(data, &bsonRepr)
	if err != nil {
		return err
	}
	*m = mainCategoryValue[bsonRepr.String]
	return nil
}

func (m Main) MarshalJSON() ([]byte, error) {
	return json.Marshal(mainCategoryName[m])
}

func (m *Main) UnmarshalJSON(data []byte) error {
	var jsonRepr string

	if err := json.Unmarshal(data, &jsonRepr); err != nil {
		return err
	}
	*m = mainCategoryValue[jsonRepr]

	return nil
}
