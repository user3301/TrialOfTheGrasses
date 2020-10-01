package gosoln

func Subsets(nums []int) [][]int {
	var ans [][]int
	subsetsBacktrack(&ans, []int{}, nums, 0)
	return ans
}

func subsetsBacktrack(ans *[][]int, subarr, nums []int, cursor int) {
	in := make([]int, len(subarr))
	copy(in, subarr)
	*ans = append(*ans, in)
	for i := cursor; i < len(nums); i++ {
		subarr = append(subarr, nums[i])
		subsetsBacktrack(ans, subarr, nums, i+1)
		subarr = subarr[: len(subarr)-1 : len(subarr)]
	}
}
