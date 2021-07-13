package test

import (
	"algo/data"
	"algo/sorting"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T){
	var length = len(data.Slice100())-1
	sorted := sorting.QuickSort(data.Slice100(),0,length)
	res1 := sort.SliceIsSorted(sorted, 
		func(i, j int) bool { 
		return sorted[i] < sorted[j] 
	})

		if res1 != true {
			t.Error("Slice is not sorted")
		}
}