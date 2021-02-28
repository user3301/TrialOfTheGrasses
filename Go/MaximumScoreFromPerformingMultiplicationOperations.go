package gosoln

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

func maximumScore(nums, multipliers []int) int {
	ans := minInt
	n := len(nums)
	m := len(multipliers)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		for j := 0; j <= i; j++ {
			var pickLeft, pickRight int
			if j == 0 {
				pickLeft = minInt
			} else {
				pickLeft = dp[j-1][i-j] + multipliers[i-1]*nums[j-1]
			}
			if j == i {
				pickRight = minInt
			} else {
				pickRight = dp[j][i-j-1] + multipliers[i-1]*nums[n-i+j]
			}
			dp[j][i-j] = max(pickLeft, pickRight)
			if i == m {
				ans = max(ans, dp[j][i-j])
			}
		}
	}
	return ans
}
