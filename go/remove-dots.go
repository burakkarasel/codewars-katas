package main

import "strings"

func main() {

}

func stringConverter(str string) string {
	return strings.ReplaceAll(str, ".", "-")
}
