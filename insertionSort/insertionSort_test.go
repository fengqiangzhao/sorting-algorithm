package insertionSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := insertionSort(sequence.Seq)
	if isSorted := sort.IsSorted(arr); !isSorted {
		t.Errorf("insertion sort failed")
	}
	fmt.Printf("%v\n", arr)
}

func TestInsertionSortII(t *testing.T) {
	arr := insertionSortII(sequence.Seq)
	if isSorted := sort.IsSorted(arr); !isSorted {
		t.Errorf("insertion sort II failed")
	}
	fmt.Printf("%v\n", arr)
}
