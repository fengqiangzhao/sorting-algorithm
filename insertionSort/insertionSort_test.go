package insertionSort

import (
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := insertionSort(sequence.Seq)
	if is_sorted := sort.IsSorted(arr); !is_sorted {
		t.Errorf("insertion sort failed")
	}
}
