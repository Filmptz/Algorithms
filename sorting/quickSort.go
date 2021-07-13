package sorting

func QuickSort(slice []int,lo int ,hi int) []int{

	var index = Partition(slice, lo, hi)

    if (lo < index - 1){
		QuickSort(slice, lo, index - 1)
	}
	 if (index < hi){
		QuickSort(slice, index, hi)
	 }
	  
	return slice
}

func swap(low *int, high *int){
	temp := *low
	*low = *high
	*high = temp
}

func Partition(slice []int,lo int,hi int) int{
	var middle = slice[(lo+hi)/2]
	
	for lo <= hi {
		

		for slice[lo] < middle {
			lo++
		}

		for slice[hi] > middle{
			hi--
		}
		if (lo <= hi) {
			swap(&slice[lo],&slice[hi])	
			lo++
			hi--

	 	 }
	}
	return lo
}