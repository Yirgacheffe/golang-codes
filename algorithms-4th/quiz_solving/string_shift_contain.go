package main

import (
	"fmt"
	"strings"
)

func find(s, t string) {
	double := func(s string) string {
		return s + s
	}

	ds := double(s)
	has := strings.Contains(ds, t)

	fmt.Println(has)
}

func main() {

	s1 := "abcd" // 'abcd' shift move -> 'abcdabcd'
	s2 := "cdab"

	find(s1, s2)
}
