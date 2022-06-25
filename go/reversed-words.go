package main

import "strings"

func ReverseWords(str string) string {
	slc := strings.Split(str, " ")
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[j], slc[i] = slc[i], slc[j]
	}
	return strings.Join(slc, " ")
}
