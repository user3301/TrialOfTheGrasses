package gosoln

func NumIdenticalPairs(nums []int) int {
	m := make(map[int][]int)
	for k, v := range nums {
		m[v] = append(m[v], k)
	}
	ans := 0
	for _, v := range m {
		if len(v) > 1 {
			ans += (len(v) - 1) * (1 + (len(v) - 1)) / 2
		}
	}
	return ans
}
