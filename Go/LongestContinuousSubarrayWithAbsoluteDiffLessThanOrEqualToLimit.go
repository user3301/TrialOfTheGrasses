package gosoln

func LongestSubarray(nums []int, limit int) int {
	maxQueue := make([]int, 0, len(nums))
	minQueue := make([]int, 0, len(nums))
	l, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		for len(minQueue) != 0 && nums[i] < minQueue[len(minQueue)-1] {
			minQueue = minQueue[:len(minQueue)-1]
		}
		for len(maxQueue) != 0 && nums[i] > maxQueue[len(maxQueue)-1] {
			maxQueue = maxQueue[:len(maxQueue)-1]
		}
		minQueue = append(minQueue, nums[i])
		maxQueue = append(maxQueue, nums[i])
		for maxQueue[0]-minQueue[0] > limit {
			if maxQueue[0] == nums[l] {
				maxQueue = maxQueue[1:]
			}
			if minQueue[0] == nums[l] {
				minQueue = minQueue[1:]
			}
			l++
		}
		ans = max(ans, i-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
