package category

type ECategory int32

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	C_CULTURE ECategory = iota + 1
	C_ENTERTAINMENT
	C_FOOD
	C_HOSPITALITY
)

type Category struct {
	Category ECategory

	Culture       Culture
	Enterteinment Entertainment
	Food          Food
	Hospitality   Hospitality
}

type builderI interface {
	Culture(v Culture) builderI
	Entertainment(v Entertainment) builderI
	Food(v Food) builderI
	Hospitality(v Hospitality) builderI

	Build() Category
}

type Builder struct {
	category ECategory

	culture       Culture
	enterteinment Entertainment
	food          Food
	hospitality   Hospitality
}

var _ builderI = Builder{}

func (b Builder) Culture(v Culture) builderI {
	b.category = C_CULTURE
	b.culture = v
	return b
}

func (b Builder) Entertainment(v Entertainment) builderI {
	b.category = C_ENTERTAINMENT
	b.enterteinment = v
	return b
}

func (b Builder) Food(v Food) builderI {
	b.category = C_FOOD
	b.food = v
	return b
}

func (b Builder) Hospitality(v Hospitality) builderI {
	b.category = C_HOSPITALITY
	b.hospitality = v
	return b
}

func (b Builder) Build() Category {
	return Category{
		Category:      b.category,
		Culture:       b.culture,
		Enterteinment: b.enterteinment,
		Food:          b.food,
		Hospitality:   b.hospitality,
	}
}

func NewBuilder() builderI {
	return Builder{}
}
