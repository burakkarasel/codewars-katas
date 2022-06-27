package kata

import "math"

func NewAvg(arr []float64, navg float64) int64 {
	var total float64
	for _, v := range arr {
		total += v
	}
	final := math.Ceil(navg*float64(len(arr)+1) - total)

	if final < 0 {
		return -1
	}
	return int64(final)
}
