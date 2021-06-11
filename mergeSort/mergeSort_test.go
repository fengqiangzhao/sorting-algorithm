package mergeSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestMergeSort(t *testing.T) {
	seq := mergeSort(sequence.Seq)
	if isSorted := sort.IsSorted(seq); !isSorted {
		t.Errorf("merge sort failed")
	}
	fmt.Printf("%v\n", seq)
}

func TestMergeSortBottomUp(t *testing.T) {
	seq := sequence.Seq
	seq.Show()
	mergeSortBottomUp(seq)
	seq.Show()
}
