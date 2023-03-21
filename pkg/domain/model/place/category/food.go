package category

type Food int32

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers
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
