package gosoln

import (
	"sort"

	"github.com/leetcode/pkg/utils"

	"github.com/leetcode/pkg/types"
)

// time complexity O(nlogn)
func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) <= 1 {
		return true
	}

	diff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// O(n) optimized solution  see https://en.wikipedia.org/wiki/Arithmetic_progression
func canMakeArithmeticProgression(arr []int) bool {
	set := make(map[int]types.Void)
	maxUint := ^uint(0)
	min := int(maxUint >> 1)
	max := -min - 1

	// find max and min in arr
	for _, v := range arr {
		max = utils.MaxInt(max, v)
		min = utils.MinInt(min, v)
		set[v] = struct{}{}
	}
	gap := (max - min) / (len(arr) - 1)
	if gap == 0 {
		return true
	}
	i := 0
	for i < len(arr) {
		if _, exist := set[min]; !exist {
			return false
		}
		min += gap
		i++
	}
	return true
}
