package category

type Culture int32

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
