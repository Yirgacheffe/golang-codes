package main

import "fmt"

func maxDivide(a, b int) int {
	for a%b == 0 {
		a /= b
	}
	return a // ~~~~~~
}

func isUgly(num int) bool {
	num = maxDivide(num, 2)
	num = maxDivide(num, 3)
	num = maxDivide(num, 5)
	return num == 1
}

func nthUglyNum(num int) int {
	i := 1
	counter := 1

	for num > counter {
		i++
		if isUgly(i) {
			counter++
		}
	}
	return i
}

func main() {
	uglyNbr := nthUglyNum(150)
	fmt.Printf("150th ugly no. is %d\n", uglyNbr)
}
