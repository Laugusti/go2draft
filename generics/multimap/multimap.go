// Package multimap provoides an implementation of map that may contain one or more values
package multimap

// mappable is a contract with two types and the first type is comparable using '=='
contract mappable(k K, _ V) {
	k == k
}

// MultiMap is map that may contain one or more values. This is a parameterized type.
type MultiMap(type K, V mappable) struct {
	m map[K][]V
}

// Get returns the list of values for the key in the MultiMap
func (m *MultiMap(K, V)) Get(k K) []V {
	return m.m[k]
}

// Set sets the key in that MultiMap to the specified value
func (m *MultiMap(K, V)) Set(k K, v V) {
	m.m[k] = []V{v}
}

// Add sets the value for the specified key is the key is not in the map or
// adds the value to the list of values for the key
func (m *MultiMap(K, V)) Add(k K, v V) {
	if len(m.Get(k)) == 0 {
		m.Set(k, v)
	} else {
		m.m[k] = append(m.m[k], v)
	}
}

// AddAll adds the list of values to map for the specified key
func (m *MultiMap(K, V)) AddAll(k K, values []V) {
	for _, v := range values {
		m.Add(k, v)
	}
}

// AsRawMap returns the MultiMap as a go map using the value selector
func (m *MultiMap(K, V)) AsRawMap(valueSelector func(K, []V) V) map[K]V {
	rm := make(map[K]V, len(m.m))
	for k, values := m.m {
		v := valueSelector(k, values)
		rm[k] = v
	}
	return rm
}

// NewMap returns a new MultiMap
func NewMap(type K, V mappable)() *MultiMap(K, V) {
	// TODO: can infer type arguments from value??
	return &MultiMap(K, V){map[K][]V{}}
}

// MapOf creates a new MultiMap using a variable number of argument values and a key function
func MapOf(type K, V mappable)(keyFunc func(int, V) K, values ...V) *MultiMap(K, V) {
	m := NewMap(K, V)()
	for i, v := range values {
		m.Add(keyFunc(i, v), v)
	}
	return v
}
