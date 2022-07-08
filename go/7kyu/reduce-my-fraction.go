package kata

func ReduceFraction(fraction [2]int) [2]int {
	var divider int
	for i := 1; i <= fraction[0] && i <= fraction[1]; i++ {
		if fraction[0]%i == 0 && fraction[1]%i == 0 {
			divider = i
		}
	}
	return [2]int{fraction[0] / divider, fraction[1] / divider}
}
