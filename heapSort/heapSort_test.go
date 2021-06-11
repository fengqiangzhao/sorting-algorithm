package heapSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := sequence.Seq
	seq := heapSort(arr)
	if isSorted := sort.IsSorted(seq); !isSorted {
		t.Errorf("heap sort failed")
	}
	fmt.Printf("%v\n", seq)

}
