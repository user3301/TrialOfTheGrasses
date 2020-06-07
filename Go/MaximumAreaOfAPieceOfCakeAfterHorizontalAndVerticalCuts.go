package gosoln

import (
	"math"
	"sort"
)

func MaxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	maxhi, maxvi := 0, 0
	horizontalCuts = append(horizontalCuts, []int{0, h}...)
	verticalCuts = append(verticalCuts, []int{0, w}...)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	i := 0
	for j := 1; j < len(horizontalCuts); j++ {
		maxhi = max(maxhi, horizontalCuts[j]-horizontalCuts[i])
		i++
	}
	i = 0
	for j := 1; j < len(verticalCuts); j++ {
		maxvi = max(maxvi, verticalCuts[j]-verticalCuts[i])
		i++
	}
	maxArea := ((maxhi) * (maxvi)) % (int(math.Pow10(9)) + 7)
	return maxArea
}
