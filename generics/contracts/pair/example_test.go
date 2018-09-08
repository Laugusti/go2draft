package pair_test

import (
	"fmt"

	"github.com/Laugusti/go2draft/generics/contracts/pair"
)

func ExamplePair() {
	// type inference
	p := pair.New(1, "one")

	fmt.Println(p.Left())
	fmt.Println(p.Right())
	// Output:
	// 1
	// one
}
