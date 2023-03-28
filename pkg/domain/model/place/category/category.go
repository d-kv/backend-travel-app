package category

// Category is a enum type.
type Category int32

func (c Category) string() string {
	return categoryName[c]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	C_CULTURE Category = iota + 1
	C_ENTERTAINMENT
	C_FOOD
	C_HOSPITALITY
)

// Enum value maps for Culture.
var (
	categoryName = map[Category]string{ //nolint:gochecknoglobals // Using global maps for enums
		C_CULTURE:       "C_CULTURE",
		C_ENTERTAINMENT: "C_ENTERTAINMENT",
		C_FOOD:          "C_FOOD",
		C_HOSPITALITY:   "C_HOSPITALITY",
	}
	categoryValue = map[string]Category{ //nolint:gochecknoglobals // Using global maps for enums
		"C_CULTURE":       C_CULTURE,
		"C_ENTERTAINMENT": C_ENTERTAINMENT,
		"C_FOOD":          C_FOOD,
		"C_HOSPITALITY":   C_HOSPITALITY,
	}
)

// Classification indicates category of a place & sub_category for each category.
//
// Use getters to check category.
type Classification struct {
	Category    Category
	SubCategory int32
}

// NewClassification creates Classification object from string value of enum.
func NewClassification(cat, subCat string) Classification {
	switch categoryValue[cat] {
	case C_CULTURE:
		return NewCulture(cultureValue[subCat])

	case C_ENTERTAINMENT:
		return NewEntertainment(entertainmentValue[subCat])

	case C_FOOD:
		return NewFood(foodValue[subCat])

	case C_HOSPITALITY:
		return NewHospitality(hospitalityValue[subCat])
	}

	return Classification{}
}

// CategoryString returns string of Category enum.
func (c Classification) CategoryString() string {
	return c.Category.string()
}

// SubCategoryString returns string of SubCategory enum.
func (c Classification) SubCategoryString() string {
	switch c.Category {
	case C_CULTURE:
		return Culture(c.SubCategory).String()

	case C_ENTERTAINMENT:
		return Entertainment(c.SubCategory).String()

	case C_FOOD:
		return Food(c.SubCategory).String()

	case C_HOSPITALITY:
		return Hospitality(c.SubCategory).String()
	}

	return ""
}

// Culture returns Culture category if specified, otherwise returns 0.
func (c Classification) Culture() Culture {
	if c.Category != C_CULTURE {
		return 0
	}
	return Culture(c.SubCategory)
}

// Entertainment returns Entertainment category if specified, otherwise returns 0.
func (c Classification) Entertainment() Entertainment {
	if c.Category != C_ENTERTAINMENT {
		return 0
	}
	return Entertainment(c.SubCategory)
}

// Food returns Food category if specified, otherwise returns 0.
func (c Classification) Food() Food {
	if c.Category != C_FOOD {
		return 0
	}
	return Food(c.SubCategory)
}

// Hospitality returns Hospitality category if specified, otherwise returns 0.
func (c Classification) Hospitality() Hospitality {
	if c.Category != C_HOSPITALITY {
		return 0
	}
	return Hospitality(c.SubCategory)
}

// NewCulture is a default ctor for Culture category.
func NewCulture(v Culture) Classification {
	return Classification{
		Category:    C_CULTURE,
		SubCategory: int32(v),
	}
}

// NewEntertainment is a default ctor for Entertainment category.
func NewEntertainment(v Entertainment) Classification {
	return Classification{
		Category:    C_ENTERTAINMENT,
		SubCategory: int32(v),
	}
}

// NewFood is a default ctor for Food category.
func NewFood(v Food) Classification {
	return Classification{
		Category:    C_FOOD,
		SubCategory: int32(v),
	}
}

// NewHospitality is a default ctor for Hospitality category.
func NewHospitality(v Hospitality) Classification {
	return Classification{
		Category:    C_HOSPITALITY,
		SubCategory: int32(v),
	}
}
