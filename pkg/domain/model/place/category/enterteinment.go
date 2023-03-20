package category

type Enterteinment int32

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	EC_UNSPECIFIED     Enterteinment = 0
	EC_NIGHTCLUB       Enterteinment = 1
	EC_AMUSEMENT_PARK  Enterteinment = 2
	EC_OPEN_MIC        Enterteinment = 3
	EC_BILLIARD_CLUB   Enterteinment = 4
	EC_WATER_PARK      Enterteinment = 5
	EC_QUEST_ROOM      Enterteinment = 6
	EC_TRAMPOLINE_PARK Enterteinment = 7
	EC_CINEMA          Enterteinment = 8
	EC_BOWLING         Enterteinment = 9
)
