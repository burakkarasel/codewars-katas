package kata

func DigitalRoot(n int) int {
	for n > 9 {
		n = n%10 + n/10
	}
	return n
}

// obsolote
/* func DigitalRoot(n int) int {
	var newN int
	for n > 0 {
		newN += n % 10
		n /= 10
		if n == 0 && newN >= 10 {
			n = newN
			newN = 0
		}
	}
	return newN
} */
