package kata

import (
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	slc := strings.Split(str, "")
	for i, v := range str {
		if unicode.IsUpper(v) == true {
			slc[i] = strings.ToLower(slc[i])
		} else if unicode.IsLower(v) {
			slc[i] = strings.ToUpper(slc[i])
		}
	}
	return strings.Join(slc, "")
}
