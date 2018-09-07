package printslice

import (
	"fmt"
)

contract stringer(x T) {
	var s string = x.String()
}

// Print writes the slice to the writer one element per line.
func Print(type T)(w io.Writer, s []T) {
	for _, v := range s {
		fmt.Fprintln(w, v)
	}
}

func PrintAsString(type T stringer)(w io.Writer, s []T) {
	for _, v := range s {
		fmt.Fprintln(w, v.String())
	}
}
