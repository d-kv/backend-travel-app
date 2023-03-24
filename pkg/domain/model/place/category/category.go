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

type builderI interface {
	Culture(v Culture) builderI
	Entertainment(v Entertainment) builderI
	Food(v Food) builderI
	Hospitality(v Hospitality) builderI

	Build() Category
}

type Builder struct {
	category ECategory

	value int32
}

var _ builderI = Builder{}

func (b Builder) Culture(v Culture) builderI {
	b.category = C_CULTURE
	b.value = int32(v)
	return b
}

func (b Builder) Entertainment(v Entertainment) builderI {
	b.category = C_ENTERTAINMENT
	b.value = int32(v)
	return b
}

func (b Builder) Food(v Food) builderI {
	b.category = C_FOOD
	b.value = int32(v)
	return b
}

func (b Builder) Hospitality(v Hospitality) builderI {
	b.category = C_HOSPITALITY
	b.value = int32(v)
	return b
}

func (b Builder) Build() Category {
	return Category(b)
}

func NewBuilder() builderI {
	return Builder{}
}
