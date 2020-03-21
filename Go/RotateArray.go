package Go

// method 1 make a new array and store the rotated result, T O(n) S O(n)
func Rotate(nums []int, k int) {
	n := len(nums)
	res := make([]int, n)

	for i, v := range nums {
		res[(i+k)%n] = v
	}
	copy(nums, res)
}

// method2 reverse array in place, T O(n) S O(1), optimal solution
func RotateInPlace(nums []int, k int) {
	n := k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, n-1)
	reverse(nums, n, len(nums)-1)
}

func reverse(nums []int, l int, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
