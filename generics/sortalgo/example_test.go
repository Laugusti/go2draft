package sortalgo_test

import (
	"fmt"

	"github.com/Laugusti/go2draft/generics/sortalgo"
)

func ExampleMergeSort() {
	s := []int{33, 23, 19, 66, 0, -8}
	sortalgo.MergeSort(s)
	fmt.Println(s)
	// Output:
	// [-8 0 19 23 33 66]
}
