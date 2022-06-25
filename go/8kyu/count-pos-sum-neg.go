package kata

func CountPositivesSumNegatives(numbers []int) []int {
	sums, count := 0, 0
	for _, v := range numbers {
		if v > 0 {
			count++
		} else {
			sums += v
		}
	}
	return []int{count, sums}
}
