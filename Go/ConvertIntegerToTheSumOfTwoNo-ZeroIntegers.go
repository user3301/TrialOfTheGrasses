package gosoln

import (
	"strconv"
	"strings"
)

func GetNoZeroIntegers(n int) []int {
	ans := make([]int, 0)

	for i := 1; i < n; i++ {
		a := strconv.Itoa(i)
		if !strings.Contains(a, "0") {
			comp := n - i
			b := strconv.Itoa(comp)
			if !strings.Contains(b, "0") {
				ans = append(ans, i)
				ans = append(ans, comp)
				return ans
			}
		}
	}
	return ans
}
