package sorting

func CountingSort(arr []int) []int{
	var countMap = make(map[int]int)
    var sortSlice = make([]int,len(arr)+1)
    
    //map to count
    for _,value := range arr {
        countMap[value]++
    }
    //sum of previous
    for i:=1; i<10;i++{
        countMap[i] += countMap[i-1] 
    }
    //sort
    for _,value := range arr{
        sortSlice[countMap[value]] = value
        countMap[value]--
    }

	return sortSlice

}