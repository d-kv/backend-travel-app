package category

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

// SubCategory is a enum type.
type SubCategory int32

type bsonSubCategory struct {
	String string
}

var _ bson.Marshaler = (*SubCategory)(nil)
var _ bson.Unmarshaler = (*SubCategory)(nil)
var _ json.Marshaler = (*SubCategory)(nil)
var _ json.Unmarshaler = (*SubCategory)(nil)

func (s SubCategory) String() string {
	return subCategoryName[s]
}

func SubCategoryFromString(scRaw string) SubCategory {
	return subCategoryValue[scRaw]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	SC_UNSPECIFIED             SubCategory = 0
	SC_RUSSIAN_CUISINE         SubCategory = 1
	SC_ITALIAN_CUISINE         SubCategory = 2
	SC_APARTMENTS              SubCategory = 3
	SC_BOWLING                 SubCategory = 4
	SC_CAMPING                 SubCategory = 5
	SC_GALLERY                 SubCategory = 6
	SC_AMUSEMENT_PARK          SubCategory = 7
	SC_ARCHITECTURAL_MONUMENTS SubCategory = 8
	SC_BEER_HOUSE              SubCategory = 9
	SC_PAB                     SubCategory = 10
	SC_VEGAN_MENU              SubCategory = 11
	SC_OPEN_MIC                SubCategory = 12
	SC_NIGHTCLUB               SubCategory = 13
	SC_COFFEE_HOUSE            SubCategory = 14
	SC_LIBRARY                 SubCategory = 15
	SC_RESORT                  SubCategory = 16
	SC_MOTEL                   SubCategory = 17
	SC_CONFECTIONERY           SubCategory = 18
	SC_JAPANESE_CUISINE        SubCategory = 19
	SC_TRAMPOLINE_PARK         SubCategory = 20
	SC_THEATRE                 SubCategory = 21
	SC_WATER_PARK              SubCategory = 22
	SC_QUEST_ROOM              SubCategory = 23
	SC_FESTIVAL                SubCategory = 24
	SC_KAFE                    SubCategory = 25
	SC_MUSEUM                  SubCategory = 26
	SC_GEORGIAN_CUISINE        SubCategory = 27
	SC_HOTEL                   SubCategory = 28
	SC_BILLIARD_CLUB           SubCategory = 29
	SC_CINEMA                  SubCategory = 30
	SC_AMERICAN_CUISINE        SubCategory = 31
	SC_BAR                     SubCategory = 32
	SC_STEAK                   SubCategory = 33
	SC_HOSTEL                  SubCategory = 34
)

var (
	subCategoryName = map[SubCategory]string{ //nolint:gochecknoglobals // Using global maps for enums
		SC_UNSPECIFIED:             "SC_UNSPECIFIED",
		SC_RUSSIAN_CUISINE:         "SC_RUSSIAN_CUISINE",
		SC_ITALIAN_CUISINE:         "SC_ITALIAN_CUISINE",
		SC_APARTMENTS:              "SC_APARTMENTS",
		SC_BOWLING:                 "SC_BOWLING",
		SC_CAMPING:                 "SC_CAMPING",
		SC_GALLERY:                 "SC_GALLERY",
		SC_AMUSEMENT_PARK:          "SC_AMUSEMENT_PARK",
		SC_ARCHITECTURAL_MONUMENTS: "SC_ARCHITECTURAL_MONUMENTS",
		SC_BEER_HOUSE:              "SC_BEER_HOUSE",
		SC_PAB:                     "SC_PAB",
		SC_VEGAN_MENU:              "SC_VEGAN_MENU",
		SC_OPEN_MIC:                "SC_OPEN_MIC",
		SC_NIGHTCLUB:               "SC_NIGHTCLUB",
		SC_COFFEE_HOUSE:            "SC_COFFEE_HOUSE",
		SC_LIBRARY:                 "SC_LIBRARY",
		SC_RESORT:                  "SC_RESORT",
		SC_MOTEL:                   "SC_MOTEL",
		SC_CONFECTIONERY:           "SC_CONFECTIONERY",
		SC_JAPANESE_CUISINE:        "SC_JAPANESE_CUISINE",
		SC_TRAMPOLINE_PARK:         "SC_TRAMPOLINE_PARK",
		SC_THEATRE:                 "SC_THEATRE",
		SC_WATER_PARK:              "SC_WATER_PARK",
		SC_QUEST_ROOM:              "SC_QUEST_ROOM",
		SC_FESTIVAL:                "SC_FESTIVAL",
		SC_KAFE:                    "SC_KAFE",
		SC_MUSEUM:                  "SC_MUSEUM",
		SC_GEORGIAN_CUISINE:        "SC_GEORGIAN_CUISINE",
		SC_HOTEL:                   "SC_HOTEL",
		SC_BILLIARD_CLUB:           "SC_BILLIARD_CLUB",
		SC_CINEMA:                  "SC_CINEMA",
		SC_AMERICAN_CUISINE:        "SC_AMERICAN_CUISINE",
		SC_BAR:                     "SC_BAR",
		SC_STEAK:                   "SC_STEAK",
		SC_HOSTEL:                  "SC_HOSTEL",
	}
	subCategoryValue = map[string]SubCategory{ //nolint:gochecknoglobals // Using global maps for enums
		"SC_UNSPECIFIED":             SC_UNSPECIFIED,
		"SC_RUSSIAN_CUISINE":         SC_RUSSIAN_CUISINE,
		"SC_ITALIAN_CUISINE":         SC_ITALIAN_CUISINE,
		"SC_APARTMENTS":              SC_APARTMENTS,
		"SC_BOWLING":                 SC_BOWLING,
		"SC_CAMPING":                 SC_CAMPING,
		"SC_GALLERY":                 SC_GALLERY,
		"SC_AMUSEMENT_PARK":          SC_AMUSEMENT_PARK,
		"SC_ARCHITECTURAL_MONUMENTS": SC_ARCHITECTURAL_MONUMENTS,
		"SC_BEER_HOUSE":              SC_BEER_HOUSE,
		"SC_PAB":                     SC_PAB,
		"SC_VEGAN_MENU":              SC_VEGAN_MENU,
		"SC_OPEN_MIC":                SC_OPEN_MIC,
		"SC_NIGHTCLUB":               SC_NIGHTCLUB,
		"SC_COFFEE_HOUSE":            SC_COFFEE_HOUSE,
		"SC_LIBRARY":                 SC_LIBRARY,
		"SC_RESORT":                  SC_RESORT,
		"SC_MOTEL":                   SC_MOTEL,
		"SC_CONFECTIONERY":           SC_CONFECTIONERY,
		"SC_JAPANESE_CUISINE":        SC_JAPANESE_CUISINE,
		"SC_TRAMPOLINE_PARK":         SC_TRAMPOLINE_PARK,
		"SC_THEATRE":                 SC_THEATRE,
		"SC_WATER_PARK":              SC_WATER_PARK,
		"SC_QUEST_ROOM":              SC_QUEST_ROOM,
		"SC_FESTIVAL":                SC_FESTIVAL,
		"SC_KAFE":                    SC_KAFE,
		"SC_MUSEUM":                  SC_MUSEUM,
		"SC_GEORGIAN_CUISINE":        SC_GEORGIAN_CUISINE,
		"SC_HOTEL":                   SC_HOTEL,
		"SC_BILLIARD_CLUB":           SC_BILLIARD_CLUB,
		"SC_CINEMA":                  SC_CINEMA,
		"SC_AMERICAN_CUISINE":        SC_AMERICAN_CUISINE,
		"SC_BAR":                     SC_BAR,
		"SC_STEAK":                   SC_STEAK,
		"SC_HOSTEL":                  SC_HOSTEL,
	}
)

func (s SubCategory) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bsonSubCategory{
		String: subCategoryName[s],
	})
}

func (s *SubCategory) UnmarshalBSON(d []byte) error {
	var bsonRepr bsonSubCategory
	err := bson.Unmarshal(d, &bsonRepr)
	if err != nil {
		return err
	}
	*s = subCategoryValue[bsonRepr.String]
	return nil
}

func (s SubCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(subCategoryName[s])
}

func (s *SubCategory) UnmarshalJSON(d []byte) error {
	var jsonRepr string

	if err := json.Unmarshal(d, &jsonRepr); err != nil {
		return err
	}
	*s = subCategoryValue[jsonRepr]
	return nil
}
