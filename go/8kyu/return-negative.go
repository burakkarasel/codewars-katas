package kata

import "math"

func MakeNegative(x int) int {
	return -int(math.Abs(float64(x)))
}

// obsolote
/* func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
} */
