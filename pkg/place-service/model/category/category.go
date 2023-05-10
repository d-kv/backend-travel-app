package category

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

// Category stores info about category in arrays of enums.
type Category struct {
	Main []Main `json:"main" bson:"main"`
	Sub  []Sub  `json:"sub" bson:"sub"`
}

type SerializedCategory struct {
	Main []string `json:"main" bson:"main"`
	Sub  []string `json:"sub" bson:"sub"`
}

var _ bson.Marshaler = (*Category)(nil)
var _ bson.Unmarshaler = (*Category)(nil)
var _ json.Marshaler = (*Category)(nil)
var _ json.Marshaler = (*Category)(nil)

// Options to build Category.
type Options func(*Category)

// WithMainCategories appends given main categories to Category.
func WithMainCategories(mc ...Main) Options {
	return func(c *Category) {
		c.Main = append(c.Main, mc...)
	}
}

// WithSubCategories appends given main categories to Category.
func WithSubCategories(sc ...Sub) Options {
	return func(c *Category) {
		c.Sub = append(c.Sub, sc...)
	}
}

// New creates a new Category applying given options.
func New(opts ...Options) *Category {
	p := &Category{
		Main: make([]Main, 0),
		Sub:  make([]Sub, 0),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func NewAnyCategory() *Category {
	return &Category{
		Main: []Main{MC_UNSPECIFIED},
		Sub:  []Sub{SC_UNSPECIFIED},
	}
}

func (c *Category) IsAnyCategory() bool {
	return len(c.Main) == 1 && c.Main[0] == MC_UNSPECIFIED &&
		len(c.Sub) == 1 && c.Sub[0] == SC_UNSPECIFIED
}

func (c Category) MarshalBSON() ([]byte, error) {
	m := make([]string, 0, len(c.Main))
	s := make([]string, 0, len(c.Sub))

	for _, cat := range c.Main {
		m = append(m, mainCategoryName[cat])
	}

	for _, cat := range c.Sub {
		s = append(s, subCategoryName[cat])
	}

	return bson.Marshal(SerializedCategory{
		Main: m,
		Sub:  s,
	})
}

func (c *Category) UnmarshalBSON(data []byte) error {
	var sc SerializedCategory
	err := bson.Unmarshal(data, &sc)
	if err != nil {
		return err
	}

	*c = *sc.toCategory()
	return nil
}

func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toSerializedCategory())
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var sc SerializedCategory
	err := json.Unmarshal(data, &sc)
	if err != nil {
		return err
	}

	*c = *sc.toCategory()
	return nil
}

func (c *SerializedCategory) toCategory() *Category {
	m := make([]Main, 0, len(c.Main))
	s := make([]Sub, 0, len(c.Sub))

	for _, cat := range c.Main {
		m = append(m, mainCategoryValue[cat])
	}

	for _, cat := range c.Sub {
		s = append(s, subCategoryValue[cat])
	}

	return &Category{
		Main: m,
		Sub:  s,
	}
}

func (c *Category) toSerializedCategory() *SerializedCategory {
	m := make([]string, 0, len(c.Main))
	s := make([]string, 0, len(c.Sub))

	for _, cat := range c.Main {
		m = append(m, mainCategoryName[cat])
	}

	for _, cat := range c.Sub {
		s = append(s, subCategoryName[cat])
	}

	return &SerializedCategory{
		Main: m,
		Sub:  s,
	}
}
