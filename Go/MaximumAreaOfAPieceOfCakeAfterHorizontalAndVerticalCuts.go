package gosoln

import (
	"math"
	"sort"
)

func MaxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	maxhi, maxvi := 0, 0
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	horizontalCuts = append([]int{0}, horizontalCuts...)
	horizontalCuts = append(horizontalCuts, h)
	i := 0
	for j := 1; j < len(horizontalCuts); j++ {
		maxhi = max(maxhi, horizontalCuts[j]-horizontalCuts[i])
		i++
	}
	verticalCuts = append([]int{0}, verticalCuts...)
	verticalCuts = append(verticalCuts, w)
	i = 0
	for j := 1; j < len(verticalCuts); j++ {
		maxvi = max(maxvi, verticalCuts[j]-verticalCuts[i])
		i++
	}
	maxArea := ((maxhi) * (maxvi)) % (int(math.Pow10(9)) + 7)
	return maxArea
}
