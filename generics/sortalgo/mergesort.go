package sortalgo

contract comparable(t T) {
	t < t
}

func MergeSort(t T comparable)(s []T) {
	if len(s) < 2 {
		return
	} else if len(s) == 2 {
		if s[0] > s[1] {
			s[0], s[1] = s[1], s[0]
		}
		return
	}

	middle := len(s)/2 + 1
	left, right := make([]T, middle), make([]T, len(s)-middle)
	copy(left, s[:middle])
	copy(right, s[middle:])
	MergeSort(left)
	MergeSort(right)

	var idx int
	for ; len(left) > 0 && len(right) > 0; idx += 1 {
		if left[0] <= right[0] {
			s[idx] = left[0]
			left = left[1:]
		} else {
			s[idx] = right[0]
			right = right[1:]
		}
	}

	s = append(s[:idx], left...)
	s = append(s[:idx], right...)
}
