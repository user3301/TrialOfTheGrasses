package gosoln

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	n := len(candidates)
	var backtrack func(temp []int, remain, idx int)
	backtrack = func(temp []int, remain, idx int) {
		if remain == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		if remain < 0 {
			return
		}
		for i := idx; i < n; i++ {
			temp = append(temp, candidates[i])
			target = target - candidates[i]
			backtrack(temp, target, i)
			temp = temp[:len(temp)-1]
			target += candidates[i]
		}
	}
	backtrack([]int{}, target, 0)
	return ans
}
