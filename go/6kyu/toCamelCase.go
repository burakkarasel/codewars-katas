package kata

import "strings"

func ToCamelCase(s string) string {
	if strings.Contains(s, "-") {
		slc := strings.Split(s, "-")
		for i := 1; i < len(slc); i++ {
			slc[i] = strings.ToUpper(string(slc[i][0])) + strings.ToLower(slc[i][1:])
		}
		return strings.Join(slc, "")
	}
	slc := strings.Split(s, "_")
	for i := 1; i < len(slc); i++ {
		slc[i] = strings.ToUpper(string(slc[i][0])) + strings.ToLower(slc[i][1:])
	}
	return strings.Join(slc, "")
}
