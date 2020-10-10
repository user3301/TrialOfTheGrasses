package gosoln

func CombinationSum3(k, n int) [][]int {
	var ans [][]int
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var temp []int
	var backtrack func(idx, sum int)
	backtrack = func(idx, sum int) {
		if len(temp) == k && sum == n {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		if len(temp) > k {
			return
		}
		for i := idx; i < len(candidates); i++ {
			sum += candidates[i]
			temp = append(temp, candidates[i])
			backtrack(i+1, sum)
			sum -= candidates[i]
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0, 0)
	return ans
}
