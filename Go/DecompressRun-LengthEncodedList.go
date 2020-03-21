package Go

func DecompressRLEList(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; i+1 < len(nums) && j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}
