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
