package test

import (
	"algo/data"
	"algo/sorting"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T){
	sorted := sorting.SelectionSort(data.Slice100())
	res1 := sort.SliceIsSorted(sorted, 
		func(i, j int) bool { 
		return sorted[i] < sorted[j] 
	})

		if res1 != true {
			t.Error("Slice is not sorted")
		}
}