package gosoln

func Shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	ptr := 0
	for i := 0; i < n; i++ {
		ans[ptr] = nums[i]
		ptr++
		ans[ptr] = nums[i+n]
		ptr++
	}
	return ans
}
