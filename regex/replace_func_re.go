package main

import (
	"fmt"
	"regexp"
)

var placeholderRE = regexp.MustCompile(`\:(owner|repo|branch)\b`)

func main() {
	value := ":repo"

	filled := placeholderRE.ReplaceAllStringFunc(value, func(m string) string {
		switch m {
		case ":repo":
			return "yirchaffee"
		case ":branch":
			return "golang-cdd"
		default:
			panic("some errors")
		}
	})

	fmt.Println(filled)
}
