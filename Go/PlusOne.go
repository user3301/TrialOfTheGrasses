package gosoln

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	//var ans []int = make([]int, len(digits)+1)
	//ans[0] =1

	ans := append([]int{1}, digits...)
	return ans
}
