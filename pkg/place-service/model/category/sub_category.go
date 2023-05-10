package category

// Sub is a enum type.
type Sub int32

func (s Sub) String() string {
	return subCategoryName[s]
}

func SubCategoryFromString(scRaw string) Sub {
	return subCategoryValue[scRaw]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	SC_UNSPECIFIED             Sub = 0
	SC_RUSSIAN_CUISINE         Sub = 1
	SC_ITALIAN_CUISINE         Sub = 2
	SC_APARTMENTS              Sub = 3
	SC_BOWLING                 Sub = 4
	SC_CAMPING                 Sub = 5
	SC_GALLERY                 Sub = 6
	SC_AMUSEMENT_PARK          Sub = 7
	SC_ARCHITECTURAL_MONUMENTS Sub = 8
	SC_BEER_HOUSE              Sub = 9
	SC_PAB                     Sub = 10
	SC_VEGAN_MENU              Sub = 11
	SC_OPEN_MIC                Sub = 12
	SC_NIGHTCLUB               Sub = 13
	SC_COFFEE_HOUSE            Sub = 14
	SC_LIBRARY                 Sub = 15
	SC_RESORT                  Sub = 16
	SC_MOTEL                   Sub = 17
	SC_CONFECTIONERY           Sub = 18
	SC_JAPANESE_CUISINE        Sub = 19
	SC_TRAMPOLINE_PARK         Sub = 20
	SC_THEATRE                 Sub = 21
	SC_WATER_PARK              Sub = 22
	SC_QUEST_ROOM              Sub = 23
	SC_FESTIVAL                Sub = 24
	SC_KAFE                    Sub = 25
	SC_MUSEUM                  Sub = 26
	SC_GEORGIAN_CUISINE        Sub = 27
	SC_HOTEL                   Sub = 28
	SC_BILLIARD_CLUB           Sub = 29
	SC_CINEMA                  Sub = 30
	SC_AMERICAN_CUISINE        Sub = 31
	SC_BAR                     Sub = 32
	SC_STEAK                   Sub = 33
	SC_HOSTEL                  Sub = 34
)

var (
	subCategoryName = map[Sub]string{ //nolint:gochecknoglobals // Using global maps for enums
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
	subCategoryValue = map[string]Sub{ //nolint:gochecknoglobals // Using global maps for enums
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
