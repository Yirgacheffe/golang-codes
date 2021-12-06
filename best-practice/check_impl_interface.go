package main

import (
	"fmt"
	"strings"
)

type Capital interface {
	Title(s string) string
}

type fetchIt string

// Compile time assertion, why ....?
var _ Capital = (*fetchIt)(nil)

func (f *fetchIt) Title(s string) string {
	return strings.ToTitle(s)
}

func main() {
	fmt.Println("check if fetchIt implement a interface with func(s string)")
}
