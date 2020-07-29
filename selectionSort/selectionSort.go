package selectionSort

import "sort_algorithm/sequence"

func selectionSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
	return arr
}
