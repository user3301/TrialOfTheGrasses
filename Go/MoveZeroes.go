package Go

func MoveZeroes(nums []int) {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
