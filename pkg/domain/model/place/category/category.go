package category

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Category is a enum type.
type Category int32

func (c Category) string() string {
	return categoryName[c]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	C_UNSPECIFIED Category = iota + 1
	C_CULTURE
	C_ENTERTAINMENT
	C_FOOD
	C_HOSPITALITY
)

// Enum value maps for Culture.
var (
	categoryName = map[Category]string{ //nolint:gochecknoglobals // Using global maps for enums
		C_UNSPECIFIED:   "C_UNSPECIFIED",
		C_CULTURE:       "C_CULTURE",
		C_ENTERTAINMENT: "C_ENTERTAINMENT",
		C_FOOD:          "C_FOOD",
		C_HOSPITALITY:   "C_HOSPITALITY",
	}
	categoryValue = map[string]Category{ //nolint:gochecknoglobals // Using global maps for enums
		"C_UNSPECIFIED":   C_UNSPECIFIED,
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
	category    Category
	subCategory int32
}

type bsonStruct struct {
	Category    string
	SubCategory string
}

func (c Classification) String() string {
	var sb strings.Builder

	sb.WriteString(c.CategoryString())
	sb.WriteString(" ")
	sb.WriteString(c.SubCategoryString())

	return sb.String()
}

func (c Classification) MarshalBSON() ([]byte, error) {
	marshalStruct := bsonStruct{
		Category:    c.CategoryString(),
		SubCategory: c.SubCategoryString(),
	}

	return bson.Marshal(marshalStruct)
}

func (c *Classification) UnmarshalBSON(data []byte) error {
	var clss bsonStruct

	if err := bson.Unmarshal(data, &clss); err != nil {
		return err
	}

	*c = NewClassification(clss.Category, clss.SubCategory)
	return nil
}

// NewClassification creates Classification object from string value of enum.
func NewClassification(cat, subCat string) Classification {
	var cl Classification
	switch categoryValue[cat] {
	case C_CULTURE:
		return NewCulture(cultureValue[subCat])

	case C_ENTERTAINMENT:
		return NewEntertainment(entertainmentValue[subCat])

	case C_FOOD:
		return NewFood(foodValue[subCat])

	case C_HOSPITALITY:
		return NewHospitality(hospitalityValue[subCat])

	case C_UNSPECIFIED:
		return cl

	default:
		return cl
	}
}

// CategoryString returns string of Category enum.
func (c Classification) CategoryString() string {
	return c.category.string()
}

// SubCategoryString returns string of SubCategory enum.
func (c Classification) SubCategoryString() string {
	switch c.category {
	case C_CULTURE:
		return Culture(c.subCategory).String()

	case C_ENTERTAINMENT:
		return Entertainment(c.subCategory).String()

	case C_FOOD:
		return Food(c.subCategory).String()

	case C_HOSPITALITY:
		return Hospitality(c.subCategory).String()

	case C_UNSPECIFIED:
		return ""

	default:
		return ""
	}
}

// Culture returns Culture category if specified, otherwise returns 0.
func (c Classification) Culture() Culture {
	if c.category != C_CULTURE {
		return 0
	}
	return Culture(c.subCategory)
}

// Entertainment returns Entertainment category if specified, otherwise returns 0.
func (c Classification) Entertainment() Entertainment {
	if c.category != C_ENTERTAINMENT {
		return 0
	}
	return Entertainment(c.subCategory)
}

// Food returns Food category if specified, otherwise returns 0.
func (c Classification) Food() Food {
	if c.category != C_FOOD {
		return 0
	}
	return Food(c.subCategory)
}

// Hospitality returns Hospitality category if specified, otherwise returns 0.
func (c Classification) Hospitality() Hospitality {
	if c.category != C_HOSPITALITY {
		return 0
	}
	return Hospitality(c.subCategory)
}

// NewCulture is a default ctor for Culture category.
func NewCulture(v Culture) Classification {
	return Classification{
		category:    C_CULTURE,
		subCategory: int32(v),
	}
}

// NewEntertainment is a default ctor for Entertainment category.
func NewEntertainment(v Entertainment) Classification {
	return Classification{
		category:    C_ENTERTAINMENT,
		subCategory: int32(v),
	}
}

// NewFood is a default ctor for Food category.
func NewFood(v Food) Classification {
	return Classification{
		category:    C_FOOD,
		subCategory: int32(v),
	}
}

// NewHospitality is a default ctor for Hospitality category.
func NewHospitality(v Hospitality) Classification {
	return Classification{
		category:    C_HOSPITALITY,
		subCategory: int32(v),
	}
}
