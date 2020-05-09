package gosoln

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	cur, sorted := 1, 0

	for cur < n {
		if nums[cur] == nums[sorted] {
			cur++
		} else {
			nums[sorted+1] = nums[cur]
			sorted++
			cur++
		}
	}
	return sorted + 1
}
