package kata

func Divisions(n, divisor int) int {
	var divisionCount int
	for n >= divisor {
		n /= divisor
		divisionCount++
	}
	return divisionCount
}
