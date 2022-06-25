package kata

import (
	"fmt"
	"strings"
)

func MultiTable(number int) (r string) {
	var slc []string
	for i := 1; i <= 10; i++ {
		val := i * number
		slc = append(slc, fmt.Sprintf("%d * %d = %d", i, number, val))
	}
	r = strings.Join(slc, "\n")
	return
}
