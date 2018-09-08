package slices_test

import (
	"fmt"

	"github.com/Laugusti/go2draft/generics/slices"
)

func ExampleMap() {
	s := []int{1, 2, 3}
	floats := slices.Map(s, func(i int) float64 { return float64(i) })

	fmt.Println(floats)
	// Output:
	// [1 2 3]
}

func ExampleReduce() {
	s := []int{1, 2, 3}
	sum := slices.Reduce(s, 0, func(i, j int) int { return i + j })

	fmt.Println(sum)
	// Output:
	// 6
}

func ExampleFilter() {
	s := []int{1, 2, 3}
	evens := slices.Filter(s, func(i int) bool { return i%2 == 0 })

	fmt.Println(evens)
	// Output:
	// [2]
}
