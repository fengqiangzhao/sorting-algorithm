package bubbleSort

import (
	"sort_algorithm/sequence"
)

func bubbleSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func bubbleSortOptimize(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		flag := true

		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		// 内循环没有改变任何元素的顺序, 即为排序成功, 不需要继续循环
		if flag {
			break
		}
	}

	return arr
}
