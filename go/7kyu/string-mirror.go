package kata

import "strings"

// VertMirror returns given strings each word reversed
func VertMirror(s string) string {
	slc := strings.Split(s, "\n")
	for i, word := range slc {
		letters := strings.Split(word, "")
		for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
			letters[i], letters[j] = letters[j], letters[i]
		}
		slc[i] = strings.Join(letters, "")
	}
	return strings.Join(slc, "\n")
}

// HorMirror returns given strings each word reorganized in reverse
func HorMirror(s string) string {
	slc := strings.Split(s, "\n")
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
	return strings.Join(slc, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}
