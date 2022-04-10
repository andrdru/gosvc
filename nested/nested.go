package nested

type (
	// Nested loader abstraction
	Nested interface {
		With(name string) Nested
		Has(name string) (exists bool)
	}

	nested map[string]struct{}
)

// NewNested .
func NewNested(with ...string) (ret Nested) {
	ret = make(nested)
	for _, w := range with {
		ret = ret.With(w)
	}

	return ret
}

// With add item to loader
func (n nested) With(name string) Nested {
	n[name] = struct{}{}
	return n
}

// Has check if item has in loader
func (n nested) Has(name string) (exists bool) {
	_, exists = n[name]
	return exists
}
