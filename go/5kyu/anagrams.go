package kata

import (
	"sort"
	"strings"
)

func Anagrams(word string, words []string) []string {
	var result []string
	wordSlc := strings.Split(word, "")
	sort.Strings(wordSlc)
	newWord := strings.Join(wordSlc, "")
	for i, w := range words {
		wSlc := strings.Split(w, "")
		sort.Strings(wSlc)
		if newWord == strings.Join(wSlc, "") {
			result = append(result, words[i])
		}
	}
	return result
}
