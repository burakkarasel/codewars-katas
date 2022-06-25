package main

import "math"

func SquareOrSquareRoot(arr []int) []int {
	var newSlc []int
	for _, v := range arr {
		tempR := math.Sqrt(float64(v))
		if tempR != math.Trunc(tempR) {
			newSlc = append(newSlc, v*v)
		} else {
			newSlc = append(newSlc, int(tempR))
		}
	}
	return newSlc
}
