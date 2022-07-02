package kata

import "math"

func CloseCompare(a, b, margin float64) int {
	var res int
	if math.Abs(a-b) <= margin {
		res = 0
	} else if a > b {
		res = 1
	} else {
		res = -1
	}
	return res
}
