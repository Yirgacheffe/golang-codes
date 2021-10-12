package main

import (
	"fmt"
)

func max(a, b int) int {
	var max = a
	if a < b {
		max = b
	}
	return max
}

func longestUniqueString(s string) int {
	var lastIdx [256]int
	for t := 0; t < 256; t++ {
		lastIdx[t] = -1
	}

	n := len(s)
	i := 0 // left
	result := 0

	for j := 0; j < n; j++ {
		// Find the last index of str[j]
		// Update i (starting index of current window)
		// as maximum of current value of i and last index plus 1
		i = max(i, lastIdx[s[j]]+1)
		lastIdx[s[j]] = j

		result = max(result, j-i+1)
	}

	return result // result of longest
}

func main() {
	n := longestUniqueString("geeksforgeeks")
	fmt.Println(n)
}
