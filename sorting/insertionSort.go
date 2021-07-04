package sorting

func InsertionSort(arr []int) []int {
	var sortSlice = make([]int, len(arr))
	
	for i := 1; i < len(arr); i++ {
		
		for j := i; j > 0; j--{
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		
	}
	sortSlice = arr
	return sortSlice
}