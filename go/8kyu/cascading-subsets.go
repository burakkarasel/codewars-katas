package kata

func EachCons(arr []int, n int) [][]int {
	var newSlc [][]int
	for i := 0; i <= len(arr)-n; i++ {
		newSlc = append(newSlc, arr[i:i+n])
	}
	return newSlc
}
