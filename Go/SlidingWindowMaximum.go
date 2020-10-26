package gosoln

func MaxSlidingWindow(nums []int, k int) []int {
	var (
		window []int
		ans []int
	)
	if nums == nil {
		return ans
	}

	for i, v := range nums {
		if len(window) > 0 && i - k + 1 > window[0] {
			window = window[1:]
		}
		for len(window) > 0 && v > nums[window[len(window) -1]] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i + 1 >= k {
			ans = append(ans, nums[window[0]])
		}
	}
	return ans
}
