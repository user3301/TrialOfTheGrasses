package gosoln

import "sort"

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
