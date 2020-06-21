package gosoln

func XorOperation(n int, start int) int {
	ans := 0
	realStart := start
	for i := 0; i < n; i++ {
		start = realStart + 2*i
		ans ^= start
	}
	return ans
}
