package main

import "fmt"

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxP := 0
	steps := 0

	for i := 0; i < length-1; {
		maxP = max(maxP, i+nums[i])
		if i == end {
			end = maxP
			steps++
		}

		i++
	}

	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 2, 3}

	fmt.Println(jump(nums))
}
