package kata

import "math"

func FindNextSquare(sq int64) int64 {
	if math.Sqrt(float64(sq)) != math.Trunc(math.Sqrt(float64(sq))) {
		return -1
	}
	return int64(math.Pow(math.Sqrt(float64(sq))+1, 2))
}
