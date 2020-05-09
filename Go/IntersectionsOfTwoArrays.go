package gosoln

func Intersection(nums1 []int, nums2 []int) []int {
	var ans []int

	set := map[int]bool{}

	for _, val := range nums1 {
		if _, ok := set[val]; !ok {
			set[val] = true
		}
	}

	for _, val := range nums2 {
		if set[val] {
			ans = append(ans, val)
			set[val] = false
		}
	}
	return ans
}
