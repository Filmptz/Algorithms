package test

import (
	"algo/sorting"
	"algo/data"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T){
	sorted := sorting.InsertionSort(data.Slice100())
	res1 := sort.SliceIsSorted(sorted, 
		func(i, j int) bool { 
		return sorted[i] < sorted[j] 
	})

		if res1 != true {
			t.Error("Slice is not sorted")
		}
}