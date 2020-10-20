package gosoln

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)
	var temp []int
	sort.Ints(candidates)
	visited := make([]bool, n)
	var backtrack func(idx, remain int)
	backtrack = func(idx, remain int) {
		if remain == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		if remain < 0 {
			return
		}
		for i := idx; i < n; i++ {
			if visited[i] || i > 0 && !visited[i-1] && candidates[i-1] == candidates[i] {
				continue
			}
			remain -= candidates[i]
			temp = append(temp, candidates[i])
			visited[i] = true
			backtrack(i+1, remain)
			visited[i] = false
			remain += candidates[i]
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0, target)
	return ans
}
