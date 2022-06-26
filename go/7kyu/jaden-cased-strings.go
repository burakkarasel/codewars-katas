package kata

// DEPRECATED
/* import "strings"

func ToJadenCase(str string) string {
  return strings.Title(str)
} */

import "strings"

func ToJadenCase(str string) string {
	slc := strings.Split(str, " ")
	for i := 0; i < len(slc); i++ {
		slc[i] = strings.ToUpper(string(slc[i][0])) + slc[i][1:]
	}
	return strings.Join(slc, " ")
}
