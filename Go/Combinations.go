package gosoln

func Combine(n, k int) [][]int {
	var ans [][]int
	var temp []int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(temp) == k {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		for i := idx; i <= n; i++ {
			temp = append(temp, i)
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(1)
	return ans
}
