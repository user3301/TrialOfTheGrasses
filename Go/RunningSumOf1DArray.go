package gosoln

func RunningSum(nums []int) []int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	ans := make([]int, len(nums))
	copy(ans, dp[1:])
	return ans
}
