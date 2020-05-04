package gosoln

func kLengthApart(nums []int, k int) bool {
	gap, i := 0, 0
	for nums[i] == 0 && i < len(nums)-1 {
		i++
	}

	for j := i; j < len(nums); j++ {
		if nums[j] == 1 && j != i {
			if gap >= k {
				gap = 0
			} else {
				return false
			}
		} else {
			gap++
		}
	}
	return true
}
