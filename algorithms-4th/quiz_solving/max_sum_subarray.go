package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Solution 2
func max_sum(a []int) int {
	n := len(a)

	mid := make([]int, n)
	mid[n-1] = a[n-1]

	fin := make([]int, n)
	fin[n-1] = a[n-1]

	for i := n - 2; i >= 0; i-- {
		mid[i] = max(a[i], a[i]+mid[i+1])
		fin[i] = max(mid[i], fin[i+1])
	}
	return fin[0]
}

// Solution 3
func max_sum2(a []int) int {
	n := len(a)

	mid := a[n-1]
	fin := a[n-1]

	for i := n - 2; i >= 0; i-- {
		mid = max(a[i], a[i]+mid)
		fin = max(mid, fin)
	}
	return fin
}

func main() {
	a := []int{1, -2, 3, 5, -3, 2} //  8
	b := []int{0, -2, 3, 5, -1, 2} //  9
	c := []int{-9, -2, -3, -5, -3} // -2

	fmt.Println(max_sum(a))
	fmt.Println(max_sum(b))
	fmt.Println(max_sum(c))

	fmt.Println(max_sum2(a))
	fmt.Println(max_sum2(b))
	fmt.Println(max_sum2(c))
}
