package main

import "strings"

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n; i++ {
		n--
		if str[i] != str[n] {
			return false
		}
	}
	return true
}
