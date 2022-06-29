package kata

func CountRedBeads(n int) int {
	if n == 1 || n == 0 {
		return 0
	}
	return (n - 1) * 2
}
