//nolint:revive, stylecheck, gochecknoglobals // Using SNAKE_CASE & global maps for enums
package category

// Hospitality is a enum type.
type Hospitality int32

func (h Hospitality) String() string {
	return hospitalityName[h]
}

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers.
const (
	HC_UNSPECIFIED Hospitality = 0
	HC_CAMPING     Hospitality = 1
	HC_RESORT      Hospitality = 2
	HC_MOTEL       Hospitality = 3
	HC_HOTEL       Hospitality = 4
	HC_HOSTEL      Hospitality = 5
	HC_APARTMENTS  Hospitality = 6
)

// Enum value maps for Hospitality.

var (
	hospitalityName = map[Hospitality]string{
		HC_UNSPECIFIED: "HC_UNSPECIFIED",
		HC_CAMPING:     "HC_CAMPING",
		HC_RESORT:      "HC_RESORT",
		HC_MOTEL:       "HC_MOTEL",
		HC_HOTEL:       "HC_HOTEL",
		HC_HOSTEL:      "HC_HOSTEL",
		HC_APARTMENTS:  "HC_APARTMENTS",
	}
	hospitalityValue = map[string]Hospitality{
		"HC_UNSPECIFIED": HC_UNSPECIFIED,
		"HC_CAMPING":     HC_CAMPING,
		"HC_RESORT":      HC_RESORT,
		"HC_MOTEL":       HC_MOTEL,
		"HC_HOTEL":       HC_HOTEL,
		"HC_HOSTEL":      HC_HOSTEL,
		"HC_APARTMENTS":  HC_APARTMENTS,
	}
)
