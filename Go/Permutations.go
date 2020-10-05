package gosoln

func Permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	var backtrack func(arr []int)
	backtrack = func(arr []int) {
		if len(arr) == n {
			ans = append(ans, append([]int{}, arr...))
		} else {
			for _, v := range nums {
				if contains(arr, v) {
					continue
				}
				arr = append(arr, v)
				backtrack(arr)
				arr = arr[:len(arr)-1]
			}
		}
	}
	backtrack([]int{})
	return ans
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
