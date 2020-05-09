package gosoln

import "strconv"

func FreqAlphabets(s string) string {
	ans := make([]rune, 0)
	start := 0

	for start < len(s) {
		if start+2 < len(s) && s[start+2] == '#' {
			num, _ := strconv.Atoi(s[start : start+2])
			ans = append(ans, rune('a'+uint8(num)-1))
			start += 3
		} else {
			num, _ := strconv.Atoi(s[start : start+1])
			ans = append(ans, rune('a'+uint8(num)-1))
			start++
		}
	}
	return string(ans)
}
