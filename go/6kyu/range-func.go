package kata

// Range func takes b (begin), e(end), s(step) arguments and creates a range
func Range(b, e, s int) []int {
	var slc []int
	for i := b; i < e && b < e; b, i = b+s, i+1 {
		slc = append(slc, b)
	}
	return slc
}

//obsolote
/* func Range(start, end, stop int) []int {
	var slc = []int{}
	if start > end {
		return slc
	}
	if start == 0 && end == 0 {
		return slc
	}
	if stop == 0 {
		for i := start; i < end; i++ {
			slc = append(slc, start)
		}
		return slc
	}
	for i := start; i < end; i += stop {
		slc = append(slc, i)
	}
	return slc
} */
