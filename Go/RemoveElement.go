package gosoln

func RemoveElement(nums []int, val int) int {
	left, right := 0, 0

	for ; right < len(nums); right++ {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	return left
}
