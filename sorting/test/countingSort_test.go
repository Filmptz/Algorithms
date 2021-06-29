package test

import (
	"algo/sorting"
	"testing"
	"sort"
)

var slice_counting = []int{1,2,3,4,5,4,3,2,1,3,4}
func TestCountingSort(t *testing.T){
	sorted := sorting.CountingSort(slice_counting)
	res1 := sort.SliceIsSorted(sorted, 
		func(i, j int) bool { 
		return sorted[i] < sorted[j] 
	})

		if res1 != true {
			t.Error("Slice is not sorted")
		}
}