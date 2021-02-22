package gosoln

import "math"

func minOperations(boxes string) []int {
	ans := make([]int, len(boxes))
	for i := range ans {
		ans[i] = calculateOperations(i, boxes)
	}
	return ans
}

func calculateOperations(pos int, boxes string) int {
	ans := 0
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '1' {
			ans += int(math.Abs(float64(i - pos)))
		}
	}
	return ans
}
