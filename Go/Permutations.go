package gosoln

func Permute(nums []int) [][]int {
	var ans [][]int
	permuteBacktrack(&ans, []int{}, nums)
	return ans
}

func permuteBacktrack(ans *[][]int, arr, nums []int) {
	if len(arr) == len(nums) {
		in := make([]int, len(arr))
		copy(in, arr)
		*ans = append(*ans, in)
	} else {
		for _, num := range nums {
			if contains(arr, num) {
				continue
			}
			arr = append(arr, num)
			permuteBacktrack(ans, arr, nums)
			arr = arr[: len(arr)-1 : len(arr)]
		}
	}
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
