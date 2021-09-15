package main

import "fmt"

func main() {
	n := 3

	fmt.Println(1 << n)
	fmt.Println(n << 1)
	fmt.Println(9 ^ (1 << n))
}
