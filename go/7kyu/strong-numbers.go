package kata

func Strong(n int) string {
	copyN := n
	total := 0

	for copyN > 0 {
		total += fact(copyN % 10)
		copyN /= 10
	}

	if total == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

func fact(num int) (f int) {
	if num == 1 || num == 0 {
		return 1
	}
	f = fact((num - 1))
	f *= num
	return
}
