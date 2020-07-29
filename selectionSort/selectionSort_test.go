package selectionSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	seq := selectionSort(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("bubble sort failed")
	}
	fmt.Printf("%v\n", seq)
}
