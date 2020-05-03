package gosoln

func MaxArea(height []int) int {
	ans, l, r := 0, 0, len(height)-1

	for l < r {
		cur := (r - l) * Min(height[l], height[r])
		ans = Max(ans, cur)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
