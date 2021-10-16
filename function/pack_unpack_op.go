package main

import "fmt"

func v1(names ...string) {
	fmt.Println(names) // pack perator to slice
}

func main() {
	var names []string = []string{"Albert", "Issac"}

	v1(
		"John", "Jane", "Daxter", "Brune",
	)

	v1(names...) // unpack operator in action, from a slice
}
