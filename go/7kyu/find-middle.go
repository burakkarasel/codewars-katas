package kata

func Gimme(a [3]int) int {
	if (a[1] > a[2] && a[2] > a[0]) || (a[0] > a[2] && a[2] > a[1]) {
		return 2
	}
	if (a[2] > a[1] && a[1] > a[0]) || (a[0] > a[1] && a[1] > a[2]) {
		return 1
	}
	return 0
}
