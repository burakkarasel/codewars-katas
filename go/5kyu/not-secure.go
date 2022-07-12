package kata

import "unicode"

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, v := range str {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}
