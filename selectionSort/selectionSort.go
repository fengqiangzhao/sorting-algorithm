package selectionSort

import "sort_algorithm/sequence"

func selectionSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr.Less(j, min) {
				min = j
			}
		}
		if min != i {
			arr.Swap(min, i)
		}
	}
	return arr
}
