package printslice_test

import (
	"fmt"
	"os"

	"github.com/Laugusti/go2draft/generics/contracts/printslice"
)

func ExamplePrint() {
	printslice.Print(int)(os.Stdout, []int{1, 2, 3})
}

type Int struct {
	i int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.i)
}

func ExamplePrintAsString() {
	printslice.PrintAsString(Int)(os.Stdout, []Int{Int{1}, Int{2}, Int{3}})
}
