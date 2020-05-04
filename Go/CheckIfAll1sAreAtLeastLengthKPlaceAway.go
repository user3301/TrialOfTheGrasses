package gosoln

func kLengthApart(nums []int, k int) bool {
	count := k

	for _, num := range nums {
		if num == 1 {
			if count < k {
				return false
			}
			count = 0
		} else {
			count++
		}
	}
	return true
}
