//nolint:revive, stylecheck, gochecknoglobals // Using SNAKE_CASE & global maps for enums
package category

// Entertainment is a enum type.
type Entertainment int32

func (e Entertainment) String() string {
	return entertainmentName[e]
}

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

// Enum value maps for Enterteinment.
var (
	entertainmentName = map[Entertainment]string{
		EC_UNSPECIFIED:     "EC_UNSPECIFIED",
		EC_NIGHTCLUB:       "EC_NIGHTCLUB",
		EC_AMUSEMENT_PARK:  "EC_AMUSEMENT_PARK",
		EC_OPEN_MIC:        "EC_OPEN_MIC",
		EC_BILLIARD_CLUB:   "EC_BILLIARD_CLUB",
		EC_WATER_PARK:      "EC_WATER_PARK",
		EC_QUEST_ROOM:      "EC_QUEST_ROOM",
		EC_TRAMPOLINE_PARK: "EC_TRAMPOLINE_PARK",
		EC_CINEMA:          "EC_CINEMA",
		EC_BOWLING:         "EC_BOWLING",
	}
	entertainmentValue = map[string]Entertainment{
		"EC_UNSPECIFIED":     EC_UNSPECIFIED,
		"EC_NIGHTCLUB":       EC_NIGHTCLUB,
		"EC_AMUSEMENT_PARK":  EC_AMUSEMENT_PARK,
		"EC_OPEN_MIC":        EC_OPEN_MIC,
		"EC_BILLIARD_CLUB":   EC_BILLIARD_CLUB,
		"EC_WATER_PARK":      EC_WATER_PARK,
		"EC_QUEST_ROOM":      EC_QUEST_ROOM,
		"EC_TRAMPOLINE_PARK": EC_TRAMPOLINE_PARK,
		"EC_CINEMA":          EC_CINEMA,
		"EC_BOWLING":         EC_BOWLING,
	}
)
