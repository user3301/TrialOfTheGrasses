package Go

func RepeatedNTimes(A []int) int {
	m := make(map[int]int)

	n := len(A) / 2
	for _, val := range A {
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
