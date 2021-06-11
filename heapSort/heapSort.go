package heapSort

import "sort_algorithm/sequence"

func heapSort(arr sequence.Sequence) sequence.Sequence {
	length := len(arr)
	buildMaxHeap(arr, length)
	for i := length - 1; i >= 0; i-- {
		swap(arr, 0, i)
		length -= 1
		heapify(arr, 0, length)
	}
	return arr
}

func buildMaxHeap(arr sequence.Sequence, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr sequence.Sequence, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < length && arr[left] > arr[largest] {
		largest = left
	}
	if right < length && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, length)
	}
}

func swap(arr sequence.Sequence, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
