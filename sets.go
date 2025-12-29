package iamerrors

// Empty is public since it is used by some internal API objects for conversions between external
// string arrays and internal sets, and conversion logic requires public types today.
type Empty struct{}

// String is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type String map[string]Empty

// NewString creates a String from a list of values.
func NewString(items ...string) String {
	ss := String{}
	ss.Insert(items...)
	return ss
}

// Insert adds items to the set.
func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s String) Has(item string) bool {
	_, contained := s[item]
	return contained
}
