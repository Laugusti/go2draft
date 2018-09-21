package sortalgo

func QuickSort(t T comparable)(s []T) {
	partition(s)
}

func partition(t T comparable)(s []T) {
	if len(s) < 2 {
		// no work needed for slice with less than 2 elements
		return
	} else if len(s) == 2 {
		// for slice of length 2, swap if out of order
		if s[0] > s[1] {
			s[0], s[1] = s[1], s[0]
		}
		return
	}
	// pivot is the beginning of the slice
	pivot := 0
	// iterate through slice, and reorder around the pivot
	for i := 1; i < len(s); i++ {
		if s[pivot] > s[i] {
			if i == pivot+1 {
				// swap element with pivot if adjacent
				s[i], s[pivot] = s[pivot], s[i]
			} else {
				// otherwise, swap element with adjcent element
				// then swap the pivot and the adjacent element
				s[i], s[pivot+1] = s[pivot+1], s[i]
				s[pivot], s[pivot+1] = s[pivot+1], s[pivot]
			}
			pivot += 1
		}
	}
	partition(s[0:pivot])  // partition left
	partition(s[pivot+1:]) // partition right
}
