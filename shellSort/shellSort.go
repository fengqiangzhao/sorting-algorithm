package shellSort

import (
	"sort_algorithm/sequence"
)

func shellSortOptimize(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)
	gap := 1

	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap >= 1 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap

			for ; j >= 0 && arr[j] > temp; j -= gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = temp
		}
		gap /= 3
	}

	return arr
}

func shellSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)
	gap := 1

	for gap < length/3 {
		gap = gap*3 + 1
	}

	for gap >= 1 {
		for i := gap; i < length; i++ {
			for j := i - gap; j >= 0 && arr.Less(j+gap, j); j -= gap {
				arr.Swap(j+gap, j)
			}
		}
		gap /= 3
	}
	return arr
}
