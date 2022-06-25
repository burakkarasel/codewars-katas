package kata

func multipleOfIndex(ints []int) (slc []int) {
	for i, v := range ints {
		if i != 0 && v%i == 0 {
			slc = append(slc, v)
		}
	}
	return
}
