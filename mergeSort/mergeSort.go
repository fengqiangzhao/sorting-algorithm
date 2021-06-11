package mergeSort

import (
	"sort_algorithm/sequence"
)

// 这种归并方法每次conquer操作都会创建数组,可以继续优化
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

	mergeLeft := mergeSort(left)
	mergeRight := mergeSort(right)

	return merge(mergeLeft, mergeRight)
}

func merge2(arr sequence.Sequence, low int, mid int, high int, aux sequence.Sequence) {
	//只有一个元素时直接返回
	if low == high || mid > high {
		return
	}

	//只有两个元素时,直接比较两个元素的值,省略下面的步骤
	if low == mid {
		if arr[low] > arr[high] {
			arr.Swap(low, high)
		}
		return
	}

	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > high {
			arr[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func mergeSortBottomUp(arr sequence.Sequence) {
	var aux sequence.Sequence
	aux = append(aux, arr...)

	n := len(arr)
	for sz := 1; sz < n; sz *= 2 { //子数组大小
		for low := 0; low < n-sz; low += 2 * sz { //lo子数组索引
			high := min(low+2*sz-1, n-1)
			mid := low + sz - 1
			merge2(arr, low, mid, high, aux)
		}
	}
}
