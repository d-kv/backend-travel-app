package category

type Hospitality int32

// Don't use iota because it is easier to keep enum in sync with api using
// explicit numbers
const (
	HC_UNSPECIFIED Hospitality = 0
	HC_CAMPING     Hospitality = 1
	HC_RESORT      Hospitality = 2
	HC_MOTEL       Hospitality = 3
	HC_HOTEL       Hospitality = 4
	HC_HOSTEL      Hospitality = 5
	HC_APARTMENTS  Hospitality = 6
)
