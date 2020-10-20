package gosoln

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	var temp []int
	visited := make([]bool, n)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		for i, v := range nums {
			if visited[i] {
				continue
			}
			if i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			temp = append(temp, v)
			visited[i] = true
			backtrack(idx + 1)
			visited[i] = false
			temp = temp[: len(temp)-1 : len(temp)]
		}
	}
	backtrack(0)
	return ans
}
