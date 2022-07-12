package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	var year int
	percent /= 100
	for p0 < p {
		p0 += int(float64(p0)*percent) + aug
		year++
	}
	return year
}
