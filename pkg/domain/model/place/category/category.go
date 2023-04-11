package category

// Category stores info about category in arrays of enums.
type Category struct {
	Main []MainCategory `json:"main" bson:"main"`
	Sub  []SubCategory  `json:"sub" bson:"sub"`
}

type bsonCategory struct {
	Main []string `json:"main" bson:"main"`
	Sub  []string `json:"sub" bson:"sub"`
}

// Options to build Category.
type Options func(*Category)

// WithMainCategories appends given main categories to Category.
func WithMainCategories(mc ...MainCategory) Options {
	return func(c *Category) {
		c.Main = append(c.Main, mc...)
	}
}

// WithSubCategories appends given main categories to Category.
func WithSubCategories(sc ...SubCategory) Options {
	return func(c *Category) {
		c.Sub = append(c.Sub, sc...)
	}
}

// New creates a new Category applying given options. If one of the fields is
// not initialized by the corresponding option, the field is set to respective
// UNSPECIFIED enum.
func New(opts ...Options) *Category {
	p := &Category{}

	for _, opt := range opts {
		opt(p)
	}

	if len(p.Main) == 0 {
		p.Main = append(p.Main, MC_UNSPECIFIED)
	}

	if len(p.Sub) == 0 {
		p.Sub = append(p.Sub, SC_UNSPECIFIED)
	}

	return p
}
