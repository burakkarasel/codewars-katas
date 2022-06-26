package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) (s string) {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)
	for _, v := range chars {
		if !strings.Contains(s, v) {
			s += v
		}
	}
	return
}
