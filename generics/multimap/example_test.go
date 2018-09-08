package multimap_test

import (
	"fmt"

	"github.com/Laugusti/go2draft/generics/multimap"
)

func ExampleMultiMap() {
	m := multimap.NewMap(int, string)()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Add(2, "dos")
	m.AddAll(3, []string{"three", "tres"})
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(2))
	fmt.Println(m.Get(3))
	// Output:
	// [one]
	// [two dos]
	// [three tres]
}

func ExampleMultiMap_MapOf() {
	// keyFunc is returns the value index as the key
	keyFunc := func(i int, s string) { return i }
	m := multimap.MapOf(int, string)(keyFunc, "zero", "one", "two")

	fmt.Println(m.Get(0))
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(2))
	// Output:
	// [zero]
	// [one]
	// [two]
}

func ExampleMultiMap_AsRawMap() {
	m := multimap.NewMap(int, string)()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Add(2, "dos")
	m.AddAll(3, []string{"three", "tres"})

	// create a map[int]string using the first value for each key in the multimap
	rm := m.AsRawMap(func(i int, a []string) { return a[0] })
	fmt.Println(rm) // prints (order not guaranteed) [1:one 2:two 3:three]
}
