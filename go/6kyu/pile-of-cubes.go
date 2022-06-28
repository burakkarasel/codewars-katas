package kata

func FindNb(m int) int {
	for i := 1; m > 0; i++ {
		m -= i * i * i
		if m == 0 {
			return i
		}
	}
	return -1
}
