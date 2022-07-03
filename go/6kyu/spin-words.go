package kata

import "strings"

func SpinWords(str string) string {
	tempSlc := strings.Split(str, " ")
	var result []string
	for _, v := range tempSlc {
		if len(v) >= 5 {
			newSlc := strings.Split(v, "")
			for i, j := 0, len(newSlc)-1; i < j; i, j = i+1, j-1 {
				newSlc[i], newSlc[j] = newSlc[j], newSlc[i]
			}
			newStr := strings.Join(newSlc, "")
			result = append(result, newStr)
		} else {
			result = append(result, v)
		}
	}
	return strings.Join(result, " ")

} // SpinWords
