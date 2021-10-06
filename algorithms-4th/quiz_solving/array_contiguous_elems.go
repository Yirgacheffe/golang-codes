package main

import "fmt"

func findLength(arr []int) int {

	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	l := len(arr)
	maxLen := 1

	for i := 0; i < l; i++ {
		mn := arr[i]
		mx := arr[i]

		for j := i + 1; j < l; j++ {
			mn = min(mn, arr[j])
			mx = max(mx, arr[j])

			// contiguous means follow j++
			if (mx - mn) == (j - i) {
				maxLen = max(maxLen, mx-mn+1)
			}
		}
	}

	return maxLen
}

func main() {
	large := []int{1, 56, 58, 57, 90, 92, 94, 93, 91, 46} // should be 5
	fmt.Printf("Length of the longest contiguous subarray is: %d\n", findLength(large))
}
