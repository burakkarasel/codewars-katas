package kata

func SquareSum(numbers []int) int {
	total := 0
	for _, item := range numbers {
		total += item * item
	}
	return total
}
