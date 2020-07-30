package mergeSort

import (
	"fmt"
	"sort_algorithm/sequence"
)

func merge(left []int, right []int) sequence.Sequence {
	var result sequence.Sequence
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	fmt.Printf("res: %v\n", result)
	return result
}

func mergeSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)

	// 递归基准条件
	if length < 2 {
		return arr
	}

	middle := length / 2

	left := arr[0:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}
