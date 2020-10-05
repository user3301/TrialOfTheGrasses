package gosoln

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	n := len(nums)
	temp := []int(nil)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		ans = append(ans, append([]int(nil), temp...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	backtrack(0)
	return ans
}
