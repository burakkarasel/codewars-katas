package main

func FakeBin(x string) string {
	b := []byte(x)
	for i, v := range b {
		if v < '5' {
			b[i] = '0'
		} else {
			b[i] = '1'
		}
	}
	return string(b)
}

// obsolote
/* import (
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	newSlc := strings.Split(x, "")
	for i, v := range newSlc {
		value, _ := strconv.Atoi(v)
		if value < 5 {
			newSlc[i] = "0"
		} else {
			newSlc[i] = "1"
		}
	}

	return strings.Join(newSlc, "")
} */
