package category

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// MainCategory is a enum type.
type MainCategory int32

func (c MainCategory) string() string {
	return mainCategoryName[c]
}

//nolint:revive, stylecheck // Using SNAKE_CASE for enums
const (
	MC_UNSPECIFIED MainCategory = iota + 1
	MC_CULTURE
	MC_ENTERTAINMENT
	MC_FOOD
	MC_HOSPITALITY
)

// Enum value maps for Culture.
var (
	mainCategoryName = map[MainCategory]string{ //nolint:gochecknoglobals // Using global maps for enums
		MC_UNSPECIFIED:   "MC_UNSPECIFIED",
		MC_CULTURE:       "MC_CULTURE",
		MC_ENTERTAINMENT: "MC_ENTERTAINMENT",
		MC_FOOD:          "MC_FOOD",
		MC_HOSPITALITY:   "MC_HOSPITALITY",
	}
	mainCategoryValue = map[string]MainCategory{ //nolint:gochecknoglobals // Using global maps for enums
		"MC_UNSPECIFIED":   MC_UNSPECIFIED,
		"MC_CULTURE":       MC_CULTURE,
		"MC_ENTERTAINMENT": MC_ENTERTAINMENT,
		"MC_FOOD":          MC_FOOD,
		"MC_HOSPITALITY":   MC_HOSPITALITY,
	}
)

// Category indicates category of a place & sub_category for each category.
type Category struct {
	mainCategory MainCategory
	subCategory  int32
}

type bsonStruct struct {
	MainCategory string `bson:"main_category"`
	SubCategory  string `bson:"sub_category"`
}

func (c Category) String() string {
	var sb strings.Builder

	sb.WriteString(c.MainCategoryString())
	sb.WriteString(" ")
	sb.WriteString(c.SubCategoryString())

	return sb.String()
}

func (c Category) MarshalBSON() ([]byte, error) {
	marshalStruct := bsonStruct{
		MainCategory: c.MainCategoryString(),
		SubCategory:  c.SubCategoryString(),
	}

	return bson.Marshal(marshalStruct)
}

func (c *Category) UnmarshalBSON(data []byte) error {
	var clss bsonStruct

	if err := bson.Unmarshal(data, &clss); err != nil {
		return err
	}

	*c = NewCategory(clss.MainCategory, clss.SubCategory)
	return nil
}

var _ bson.Marshaler = (*Category)(nil)
var _ bson.Unmarshaler = (*Category)(nil)

// NewCategory creates Category object from string value of enum.
func NewCategory(mainCat, subCat string) Category {
	var cl Category
	switch mainCategoryValue[mainCat] {
	case MC_CULTURE:
		return NewCulture(cultureValue[subCat])

	case MC_ENTERTAINMENT:
		return NewEntertainment(entertainmentValue[subCat])

	case MC_FOOD:
		return NewFood(foodValue[subCat])

	case MC_HOSPITALITY:
		return NewHospitality(hospitalityValue[subCat])

	case MC_UNSPECIFIED:
		return cl

	default:
		return cl
	}
}

// MainCategoryString returns string of Category enum.
func (c Category) MainCategoryString() string {
	return c.mainCategory.string()
}

// SubCategoryString returns string of SubCategory enum.
func (c Category) SubCategoryString() string {
	switch c.mainCategory {
	case MC_CULTURE:
		return Culture(c.subCategory).String()

	case MC_ENTERTAINMENT:
		return Entertainment(c.subCategory).String()

	case MC_FOOD:
		return Food(c.subCategory).String()

	case MC_HOSPITALITY:
		return Hospitality(c.subCategory).String()

	case MC_UNSPECIFIED:
		return ""

	default:
		return ""
	}
}

// NewCulture is a default ctor for Culture category.
func NewCulture(v Culture) Category {
	return Category{
		mainCategory: MC_CULTURE,
		subCategory:  int32(v),
	}
}

// NewEntertainment is a default ctor for Entertainment category.
func NewEntertainment(v Entertainment) Category {
	return Category{
		mainCategory: MC_ENTERTAINMENT,
		subCategory:  int32(v),
	}
}

// NewFood is a default ctor for Food category.
func NewFood(v Food) Category {
	return Category{
		mainCategory: MC_FOOD,
		subCategory:  int32(v),
	}
}

// NewHospitality is a default ctor for Hospitality category.
func NewHospitality(v Hospitality) Category {
	return Category{
		mainCategory: MC_HOSPITALITY,
		subCategory:  int32(v),
	}
}
