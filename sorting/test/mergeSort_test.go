package test

import (
	"algo/data"
	"algo/sorting"
	"sort"
	"testing"
)

func TestMergerSort(t *testing.T){
	sorted := sorting.MergeSort(data.Slice100())
	res1 := sort.SliceIsSorted(sorted, 
		func(i, j int) bool { 
		return sorted[i] < sorted[j] 
	})

		if res1 != true {
			t.Error("Slice is not sorted")
		}
}