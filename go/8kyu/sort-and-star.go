package kata

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	slc := strings.Split(arr[0], "")
	return strings.Join(slc, "***")
}
