package sorting

func BubbleSort(arr []int) []int {
	var sortSlice = make([]int, len(arr))
	var isSwapped = true

	for isSwapped {
		isSwapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				isSwapped = true
			}
		}
	}

	sortSlice = arr
	return sortSlice
}