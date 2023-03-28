package category

type ECategory int32

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	C_CULTURE ECategory = iota + 1
	C_ENTERTAINMENT
	C_FOOD
	C_HOSPITALITY
)

// Category indicates category of a place & sub_category for each category.
//
// Use getters to check category.
type Category struct {
	Category ECategory `bson:"category"`

	Value int32 `bson:"sub_category"`
}

// Culture returns Culture category if specified, otherwise returns 0.
func (c Category) Culture() Culture {
	if c.Category != C_CULTURE {
		return 0
	}
	return Culture(c.Value)
}

// Entertainment returns Entertainment category if specified, otherwise returns 0.
func (c Category) Entertainment() Entertainment {
	if c.Category != C_ENTERTAINMENT {
		return 0
	}
	return Entertainment(c.Value)
}

// Food returns Food category if specified, otherwise returns 0.
func (c Category) Food() Food {
	if c.Category != C_FOOD {
		return 0
	}
	return Food(c.Value)
}

// Hospitality returns Hospitality category if specified, otherwise returns 0.
func (c Category) Hospitality() Hospitality {
	if c.Category != C_HOSPITALITY {
		return 0
	}
	return Hospitality(c.Value)
}

// NewCulture is a default ctor for Culture category.
func NewCulture(v Culture) Category {
	return Category{
		Category: C_CULTURE,
		Value:    int32(v),
	}
}

// NewEntertainment is a default ctor for Entertainment category.
func NewEntertainment(v Entertainment) Category {
	return Category{
		Category: C_ENTERTAINMENT,
		Value:    int32(v),
	}
}

// NewFood is a default ctor for Food category.
func NewFood(v Food) Category {
	return Category{
		Category: C_FOOD,
		Value:    int32(v),
	}
}

// NewHospitality is a default ctor for Hospitality category.
func NewHospitality(v Hospitality) Category {
	return Category{
		Category: C_HOSPITALITY,
		Value:    int32(v),
	}
}
