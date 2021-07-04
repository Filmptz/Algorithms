package sorting

func MergeSort(arr []int) []int {	
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr)/2	
	return MergeArr(MergeSort(arr[:middle]),MergeSort(arr[middle:]))
}

func MergeArr(leftArr []int,rightArr []int) []int{
	var lenSlice = len(leftArr)+len(rightArr)
	var SortSlice = make([]int,lenSlice)

	var lCount, rCount = 0,0
	for x:=0; x<lenSlice; x++{
		if lCount > len(leftArr)-1 && rCount <= len(rightArr)-1{
			SortSlice[x] = rightArr[rCount]
			rCount++
		}else if rCount > len(rightArr)-1 && lCount <= len(leftArr)-1 {
			SortSlice[x] = leftArr[lCount]
			lCount++
		}else if leftArr[lCount] < rightArr[rCount]{
			SortSlice[x] = leftArr[lCount]
			lCount++
		}else{
			SortSlice[x] = rightArr[rCount]
			rCount++
		}
		
	}
	return SortSlice
}