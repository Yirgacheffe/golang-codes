package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, e := regexp.Compile(`\w{3}`)
	if e != nil {
		panic(e)
	}

	fmt.Println(string(re.Find([]byte("abc aaaa abcdef aaa"))))
}
