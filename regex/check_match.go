package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	}

	if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number!")
	} else {
		fmt.Println("Not Number!")
	}

	fmt.Println(regexp.Match(`\d`, []byte("1 12 23")))
	fmt.Println(regexp.Match(`\d\d`, []byte("1 12 23")))
	fmt.Println(regexp.Match(`\w+`, []byte("1 12 23")))

	isJSON, _ := regexp.MatchString(`[/+]json[;|$]`, "application/json; charset=utf8")
	fmt.Println(isJSON)

}
