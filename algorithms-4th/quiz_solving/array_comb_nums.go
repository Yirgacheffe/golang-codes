package main

import "fmt"

func combinate(candidates []int, target int) (ans [][]int) {

	comb := []int{}
	var dfs func(target, idx int)

	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		dfs(target, idx+1)

		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1] // ??? ---- ... :_]
		}
	}

	dfs(target, 0)
	return // ----------------------------------------------
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinate(candidates, target))
}
