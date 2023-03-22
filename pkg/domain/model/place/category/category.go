package category

type ECategory int32

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
