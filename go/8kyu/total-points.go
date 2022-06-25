package kata

import "strings"

func Points(games []string) (r int) {
	for _, g := range games {
		scores := strings.Split(g, ":")
		f, s := scores[0], scores[1]
		if f > s {
			r += 3
		}
		if f == s {
			r++
		}
	}
	return
}
