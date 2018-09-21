package sortalgo

import (
	"math/rand"
	"testing"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func randomIntSlice(n, size int) []int {
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(n)
	}
	return s
}

func isSortedAsc(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
		{[]int{2, 1}},
		{[]int{3, 2, 1}},
		{randomIntSlice(100, 5)},
		{randomIntSlice(100000, 1000)},
	}

	for _, test := range tests {
		MergeSort(test.input)
		if !isSortedAsc(test.input) {
			t.Errorf("input not sorted: %v", test.input)
		}
	}
}
