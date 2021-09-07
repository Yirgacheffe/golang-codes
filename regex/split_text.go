package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := regexp.MustCompile(`a`).Split("abababacafagadatarat", 7)
	fmt.Println(s)
}
