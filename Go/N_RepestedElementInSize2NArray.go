package Go

func RepeatedNTimes(a []int) int {
	m := make(map[int]int)

	n := len(a) / 2
	for _, val := range a {
		if _, ok := m[val]; ok {
			m[val]++
			if m[val] >= n {
				return val
			}
		} else {
			m[val] = 1
		}
	}
	panic("not found")
}
