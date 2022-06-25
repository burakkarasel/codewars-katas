package kata

import (
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))
}

// obsolote
/* func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if char >= 97 && char <= 122 {
			return false
		}
	}
	return true
} */
