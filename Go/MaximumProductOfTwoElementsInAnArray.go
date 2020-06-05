package gosoln

func MaxProduct(nums []int) int {
	minInt := -int(^uint(0) >> 1)
	first, second := minInt, minInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > first {
			second = first
			first = nums[i]
		} else if nums[i] > second {
			second = nums[i]
		}
	}
	return (first - 1) * (second - 1)
}
