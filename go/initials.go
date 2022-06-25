package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) (r string) {
	slc := strings.Split(name, " ")
	r = fmt.Sprintf("%s.%s", strings.ToUpper(string(slc[0][0])), strings.ToUpper(string(slc[1][0])))
	return
}
