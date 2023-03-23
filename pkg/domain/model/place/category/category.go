package category

type ECategory int32

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	C_CULTURE ECategory = iota + 1
	C_ENTARTAINMENT
	C_FOOD
	C_HOSPITALITY
)

type Category struct {
	Category ECategory

	Culture       Culture
	Enterteinment Enterteinment
	Food          Food
	Hospitality   Hospitality
}
