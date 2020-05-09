package gosoln

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if idx, ok := m[complement]; ok {
			ans := []int{i, idx}
			return ans
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
