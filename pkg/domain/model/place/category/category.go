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
	category ECategory

	value int32
}

func (c Category) Culture() Culture {
	if c.category != C_CULTURE {
		return 0
	}
	return Culture(c.value)
}

func (c Category) Entertainment() Entertainment {
	if c.category != C_ENTERTAINMENT {
		return 0
	}
	return Entertainment(c.value)
}

func (c Category) Food() Food {
	if c.category != C_FOOD {
		return 0
	}
	return Food(c.value)
}

func (c Category) Hospitality() Hospitality {
	if c.category != C_HOSPITALITY {
		return 0
	}
	return Hospitality(c.value)
}

func NewCulture(v Culture) *Category {
	return &Category{
		category: C_CULTURE,
		value:    int32(v),
	}
}

func NewEntertainment(v Entertainment) *Category {
	return &Category{
		category: C_ENTERTAINMENT,
		value:    int32(v),
	}
}

func NewFood(v Food) *Category {
	return &Category{
		category: C_FOOD,
		value:    int32(v),
	}
}

func NewHospitality(v Hospitality) *Category {
	return &Category{
		category: C_HOSPITALITY,
		value:    int32(v),
	}
}
