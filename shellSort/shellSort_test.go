package shellSort

import (
	"sort"
	"sort_algorithm/sequence"
	"testing"
)

func TestShellSortOptimize(t *testing.T) {
	a := sequence.Seq
	a.Show()

	sortedArr := shellSortOptimize(a)
	sortedArr.Show()

	if isSorted := sort.IsSorted(sortedArr); !isSorted {
		t.Errorf("shell sort failed")
	}
}

func TestShellSort(t *testing.T) {
	a := sequence.Seq
	a.Show()

	sortedArr := shellSort(a)
	sortedArr.Show()

	if isSorted := sort.IsSorted(sortedArr); !isSorted {
		t.Errorf("shell sort failed")
	}
}
