package gosoln

func LengthOfLongestSubstring(s string) int {
	var ans int
	m := make(map[uint8]int)

	count, l, r := 0, 0, 0

	for r < len(s) {
		m[s[r]]++

		if m[s[r]] > 1 {
			count++
		}

		for count > 0 {
			if m[s[l]] > 1 {
				count--
			}
			m[s[l]] -= 1
			l++
		}
		r++
		ans = Max(ans, r-l)
	}
	return ans
}
