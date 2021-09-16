package main

import "fmt"

func median(a, b []int) float64 {

	n := len(a)
	m := len(b)

	m1 := -1
	m2 := -1

	i := 0
	j := 0

	if (m+n)%2 == 1 {
		for cnt := 0; cnt <= (m+n)/2; cnt++ {
			if i != n && j != m {
				if a[i] > b[j] {
					m1 = b[j]
					j++
				} else {
					m1 = a[i]
					i++
				}
			} else if i < n {
				m1 = a[i]
				i++
			} else {
				m1 = b[j]
				j++
			}
		}
		return float64(m1)
	} else {
		for cnt := 0; cnt <= (m+n)/2; cnt++ {
			m2 = m1
			if i != n && j != m {
				if a[i] > b[j] {
					m1 = b[j]
					j++
				} else {
					m1 = a[i]
					i++
				}
			} else if i < n {
				m1 = a[i]
				i++
			} else {
				m1 = b[j]
				j++
			}
		}
		return float64(m1+m2) / 2
	}
}

func main() {
	// Solution 1 is merge -> sort -> find
	a := []int{2, 3, 7, 20}
	b := []int{4, 5, 6, 10}
	fmt.Println(median(a, b))
}
