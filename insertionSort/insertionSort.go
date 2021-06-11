package insertionSort

import "sort_algorithm/sequence"

// 优化版
func insertionSort(arr sequence.Sequence) sequence.Sequence {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]

		//内循环中将较大的元素都向右移动,最后交换
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

// 初始版
func insertionSortII(arr sequence.Sequence) sequence.Sequence {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr.Less(j, j-1); j-- {
			arr.Swap(j, j-1)
		}
	}
	return arr
}
