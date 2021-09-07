package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`r[ua]n`)

	fmt.Printf("%q\n", re.FindAll([]byte("ran run run abc def abcd"), -1))
	fmt.Println(re.FindAllIndex([]byte("ran run run abc def abcd"), -1))
}
