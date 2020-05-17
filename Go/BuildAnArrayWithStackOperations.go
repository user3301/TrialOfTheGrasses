package gosoln

func BuildArray(target []int, n int) []string {
	ans := make([]string, 0, len(target))
	cur := 1

	for i := 0; i < len(target); i++ {
		if target[i] > cur {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
			cur++
			i--
		} else if target[i] == cur {
			ans = append(ans, "Push")
			cur++
		} else {
			panic("not possible")
		}
	}
	return ans
}
