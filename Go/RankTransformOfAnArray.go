package gosoln

import "sort"

func ArrayRankTransform(arr []int) []int {
	a := make([]int, len(arr))
	rank := map[int]int{}
	copy(a, arr)
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		if _, exist := rank[a[i]]; !exist {
			rank[a[i]] = len(rank) + 1
		}
	}
	var ans []int
	for _, v := range arr {
		ans = append(ans, rank[v])
	}
	return ans
}
