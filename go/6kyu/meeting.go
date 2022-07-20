package kata

import (
	"sort"
	"strings"
)

func Meeting(s string) string {
	slc := strings.Split(s, ";")
	var strSlc []string
	for _, v := range slc {
		tempSlc := strings.Split(v, ":")
		tempSlc[0], tempSlc[1] = tempSlc[1], tempSlc[0]
		strSlc = append(strSlc, "("+strings.ToUpper(strings.Join(tempSlc, ", "))+")")
	}
	sort.Strings(strSlc)
	return strings.Join(strSlc, "")
}
