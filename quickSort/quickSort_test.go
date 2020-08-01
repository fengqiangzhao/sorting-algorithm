package quickSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestQuickSort(t *testing.T) {
	seq := quickSort(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("quick sort failed")
	}
	fmt.Printf("%v\n", seq)
}
