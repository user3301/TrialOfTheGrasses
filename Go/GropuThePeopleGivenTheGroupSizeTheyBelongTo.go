package Go

func GroupThePeople(groupSizes []int) [][]int {
	ans := make([][]int, 0)

	m := make(map[int][]int)

	for i, v := range groupSizes {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = append(m[v], i)
		}
		if len(m[v]) == v {
			ans = append(ans, m[v])
			m[v] = []int{}
		}
	}
	return ans
}
