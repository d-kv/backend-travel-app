//nolint:revive, stylecheck // Using SNAKE_CASE for enums
package category

type Entertainment int32

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	EC_UNSPECIFIED     Entertainment = 0
	EC_NIGHTCLUB       Entertainment = 1
	EC_AMUSEMENT_PARK  Entertainment = 2
	EC_OPEN_MIC        Entertainment = 3
	EC_BILLIARD_CLUB   Entertainment = 4
	EC_WATER_PARK      Entertainment = 5
	EC_QUEST_ROOM      Entertainment = 6
	EC_TRAMPOLINE_PARK Entertainment = 7
	EC_CINEMA          Entertainment = 8
	EC_BOWLING         Entertainment = 9
)
