package bubbleSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	seq := bubbleSort(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("bubble sort failed")
	}
	fmt.Printf("%v\n", seq)
}

func Test_bubbleSortOptimize(t *testing.T) {
	seq := bubbleSortOptimize(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("optimize bubble sort failed")
	}
	fmt.Printf("%v\n", seq)
}
