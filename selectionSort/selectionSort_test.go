package selectionSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	seq := selectionSort(sequence.Seq)
	if isSorted := sort.IsSorted(seq); !isSorted {
		t.Errorf("selection sort failed")
	}
	fmt.Printf("%v\n", seq)
}
