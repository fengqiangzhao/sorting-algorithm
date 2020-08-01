package quickSort

import "sort_algorithm/sequence"

func quickSort(arr sequence.Sequence) sequence.Sequence {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr sequence.Sequence, left, right int) sequence.Sequence {
	if left < right {
		partitionIndex := partition3(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

// 分组, 把数组第一位设为pivot
func partition(arr sequence.Sequence, left, right int) int {
	pivot := left
	index := pivot

	for i := pivot + 1; i <= right; i++ {
		if arr[i] < arr[pivot] {
			index += 1
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
	arr[pivot], arr[index] = arr[index], arr[pivot]
	return index
}

// 分组, 把数组最后一位设为pivot
func partition2(arr sequence.Sequence, left, right int) int {
	pivot := right
	index := left

	for i := left; i <= right-1; i++ {
		if arr[i] <= arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[pivot], arr[index] = arr[index], arr[pivot]
	return index
}

// 双指针两头遍历数组
func partition3(arr sequence.Sequence, left, right int) int {
	pivot := left
	i := left
	j := right

	for i != j {
		for ; arr[j] >= arr[pivot] && i < j; j-- {continue}

		for ; arr[i] <= arr[pivot] && i < j; i++ {continue}

		arr[i], arr[j] = arr[j], arr[i]

	}
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}
