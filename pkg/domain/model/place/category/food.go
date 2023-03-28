//nolint:revive, stylecheck, gochecknoglobals // Using SNAKE_CASE & global maps for enums
package category

// Food is a enum type.
type Food int32

func (f Food) String() string {
	return foodName[f]
}

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	FC_UNSPECIFIED      Food = 0
	FC_BEER_HOUSE       Food = 1
	FC_CONFECTIONERY    Food = 2
	FC_COFFEE_HOUSE     Food = 3
	FC_RUSSIAN_CUISINE  Food = 4
	FC_BAR              Food = 5
	FC_ITALIAN_CUISINE  Food = 6
	FC_JAPANESE_CUISINE Food = 7
	FC_STEAK            Food = 8
	FC_AMERICAN_CUISINE Food = 9
	FC_VEGAN_MENU       Food = 10
	FC_PAB              Food = 11
	FC_GEORGIAN_CUISINE Food = 12
	FC_KAFE             Food = 13
)

// Enum value maps for Food.
var (
	foodName = map[Food]string{
		FC_UNSPECIFIED:      "FC_UNSPECIFIED",
		FC_BEER_HOUSE:       "FC_BEER_HOUSE",
		FC_CONFECTIONERY:    "FC_CONFECTIONERY",
		FC_COFFEE_HOUSE:     "FC_COFFEE_HOUSE",
		FC_RUSSIAN_CUISINE:  "FC_RUSSIAN_CUISINE",
		FC_BAR:              "FC_BAR",
		FC_ITALIAN_CUISINE:  "FC_ITALIAN_CUISINE",
		FC_JAPANESE_CUISINE: "FC_JAPANESE_CUISINE",
		FC_STEAK:            "FC_STEAK",
		FC_AMERICAN_CUISINE: "FC_AMERICAN_CUISINE",
		FC_VEGAN_MENU:       "FC_VEGAN_MENU",
		FC_PAB:              "FC_PAB",
		FC_GEORGIAN_CUISINE: "FC_GEORGIAN_CUISINE",
		FC_KAFE:             "FC_KAFE",
	}
	foodValue = map[string]Food{
		"FC_UNSPECIFIED":      FC_UNSPECIFIED,
		"FC_BEER_HOUSE":       FC_BEER_HOUSE,
		"FC_CONFECTIONERY":    FC_CONFECTIONERY,
		"FC_COFFEE_HOUSE":     FC_COFFEE_HOUSE,
		"FC_RUSSIAN_CUISINE":  FC_RUSSIAN_CUISINE,
		"FC_BAR":              FC_BAR,
		"FC_ITALIAN_CUISINE":  FC_ITALIAN_CUISINE,
		"FC_JAPANESE_CUISINE": FC_JAPANESE_CUISINE,
		"FC_STEAK":            FC_STEAK,
		"FC_AMERICAN_CUISINE": FC_AMERICAN_CUISINE,
		"FC_VEGAN_MENU":       FC_VEGAN_MENU,
		"FC_PAB":              FC_PAB,
		"FC_GEORGIAN_CUISINE": FC_GEORGIAN_CUISINE,
		"FC_KAFE":             FC_KAFE,
	}
)
