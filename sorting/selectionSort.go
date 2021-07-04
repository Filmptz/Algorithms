package sorting

func SelectionSort(arr []int) []int {
	var sortSlice = make([]int, len(arr))
	var minValue,minIndex int
	for i := 0; i < len(arr); i++ {
		minIndex = i
		for j := i+1; j <len(arr); j++{
			if arr[j] < arr[minIndex] {
				minIndex = j
				minValue = arr[j]
			}
		}
		arr[i] = minValue
		
	}
	sortSlice = arr
	
	return sortSlice
}