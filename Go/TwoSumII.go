package Go

func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}
			right--
		} else {
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			left++
		}
	}
	panic("not found")
}
