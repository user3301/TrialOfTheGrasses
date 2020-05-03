package gosoln

func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		switch {
		case numbers[left]+numbers[right] == target:
			return []int{left + 1, right + 1}
		case numbers[left]+numbers[right] > target:
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}
			right--
		default:
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			left++
		}
	}
	panic("not found")
}
