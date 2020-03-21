package Go

func SumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}

	ans := make([]int, 0, n)

	if n%2 == 1 {
		ans = append(ans, 0)
	}

	for i := 1; i < n/2+1; i++ {
		ans = append(ans, 1*i)
		ans = append(ans, -1*i)
	}
	return ans
}
