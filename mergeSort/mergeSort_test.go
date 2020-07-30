package mergeSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestMergeSort(t *testing.T) {
	seq := mergeSort(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("merge sort failed")
	}
	fmt.Printf("%v\n", seq)
}
