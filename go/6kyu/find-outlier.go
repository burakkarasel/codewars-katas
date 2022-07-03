package kata

func FindOutlier(integers []int) int {
	var evens, odds []int
	var r int
	for _, v := range integers {
		if v%2 == 0 {
			evens = append(evens, v)
		}
		if v%2 == 1 || v%2 == -1 {
			odds = append(odds, v)
		}
		if len(evens) >= 2 && len(odds) == 1 {
			r = odds[0]
			break
		}
		if len(odds) >= 2 && len(evens) == 1 {
			r = evens[0]
			break
		}
	}
	return r
}
