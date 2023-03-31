//nolint:revive, stylecheck, gochecknoglobals // Using SNAKE_CASE & global maps for enums
package category

// Culture is a enum type.
type Culture int32

func (c Culture) String() string {
	return cultureName[c]
}

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	CC_UNSPECIFIED             Culture = 0
	CC_MUSEUM                  Culture = 1
	CC_FESTIVAL                Culture = 2
	CC_GALLERY                 Culture = 3
	CC_LIBRARY                 Culture = 4
	CC_ARCHITECTURAL_MONUMENTS Culture = 5
	CC_THEATRE                 Culture = 6
)

// Enum value maps for Culture.
var (
	cultureName = map[Culture]string{
		CC_UNSPECIFIED:             "CC_UNSPECIFIED",
		CC_MUSEUM:                  "CC_MUSEUM",
		CC_FESTIVAL:                "CC_FESTIVAL",
		CC_GALLERY:                 "CC_GALLERY",
		CC_LIBRARY:                 "CC_LIBRARY",
		CC_ARCHITECTURAL_MONUMENTS: "CC_ARCHITECTURAL_MONUMENTS",
		CC_THEATRE:                 "CC_THEATRE",
	}
	cultureValue = map[string]Culture{
		"CC_UNSPECIFIED":             CC_UNSPECIFIED,
		"CC_MUSEUM":                  CC_MUSEUM,
		"CC_FESTIVAL":                CC_FESTIVAL,
		"CC_GALLERY":                 CC_GALLERY,
		"CC_LIBRARY":                 CC_LIBRARY,
		"CC_ARCHITECTURAL_MONUMENTS": CC_ARCHITECTURAL_MONUMENTS,
		"CC_THEATRE":                 CC_THEATRE,
	}
)
