package gosoln

func Intersect(nums1 []int, nums2 []int) []int {
	var ans []int

	m := map[int]int{}

	for _, v := range nums1 {
		m[v] += 1
	}

	for _, v := range nums2 {
		ans = append(ans, v)
		m[v]--
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return ans
}
