package kata

func ReverseSeq(n int) (slc []int) {
	for i := n; i > 0; i-- {
		slc = append(slc, i)
	}
	return
}
