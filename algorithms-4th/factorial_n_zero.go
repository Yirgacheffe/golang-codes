package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	temp := 1 // 0!
	n := 18
	for i := 1; i <= n; i++ {
		temp *= i
	}

	//
	xemp := factorial(18)
	if temp != xemp {
		fmt.Printf("got: %d, want: %d", xemp, temp)
	}

	// calculate 6402373705728000
	ret := 0
	for i := 1; i <= n; i++ {
		j := i
		for j%5 == 0 {
			ret++
			j /= 5
		}
	}
	fmt.Println("zero of N! is:", ret)

	// lowest 1 position
	low := 0
	for n > 0 {
		n >>= 1
		fmt.Println(n)
		low += n
	}
	fmt.Println("lowest 1 of N! is:", low)

}
