package shellSort

import (
	"fmt"
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestShellSort(t *testing.T) {
	seq := shellSort(sequence.Seq)
	if is_sorted := sort.IsSorted(seq); !is_sorted {
		t.Errorf("shell sort failed")
	}
	fmt.Printf("%v\n", seq)
}
