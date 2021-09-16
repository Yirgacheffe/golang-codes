package main

import "fmt"

func shift(s string, k int) string {
	c := []byte(s)
	n := len(c)
	t := s[n-k:]

	for i := n - 1; i >= 0; i-- {
		if i >= k {
			c[i] = c[i-k]
		} else {
			c[i] = t[i]
		}
	}

	return string(c)
}

func main() {
	s := "abcdef123456"
	fmt.Println(shift(s, 4))

	// #Try this
	// abcd1234 -> dcba4321 -> 1234abcd
}
