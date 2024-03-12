package fieldy

func New(name, value string) Field {
	return Field{name, value}
}

// Field is a struct that represents a field in a form
// - properties:
// > inmutable
type Field struct {
	name  string
	value string
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Value() string {
	return f.value
}

type Fields []Field