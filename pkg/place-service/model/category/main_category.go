package category

// Main is a enum type.
type Main int32

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
