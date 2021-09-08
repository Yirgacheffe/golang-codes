package main

import "fmt"

// solution 1 -> v % 2
// solution 2 -> v & 0x01

// solution 4 -> [256]countTable

// solution 3
func main() {
	n := 0b01001111
	num := 0

	for n > 0 {
		n &= (n - 1)
		fmt.Printf("%#b\n", n)
		num++
	}

	fmt.Println(num)
}
