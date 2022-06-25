package kata

func HowManyDalmatians(n int) (result string) {
	r := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}
	switch {
	case n <= 10:
		result = r[0]
	case n <= 50:
		result = r[1]
	case n < 101:
		result = r[2]
	case n == 101:
		result = r[3]
	}
	return
}
