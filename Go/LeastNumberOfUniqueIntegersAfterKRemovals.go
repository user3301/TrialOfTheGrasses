package gosoln

import "sort"

type NumCount struct {
	S [][]int
}

func (n NumCount) Len() int {
	return len(n.S)
}

func (n NumCount) Swap(i, j int) {
	n.S[i], n.S[j] = n.S[j], n.S[i]
}

func (n NumCount) Less(i, j int) bool {
	return n.S[i][1] < n.S[j][1]
}

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	dist := make(map[int]int)
	for _, v := range arr {
		dist[v] += 1
	}
	nc := NumCount{}
	for k, v := range dist {
		nc.S = append(nc.S, []int{k, v})
	}
	sort.Stable(nc)
	ans := len(nc.S)
	i := 0
	for k > 0 {
		if k >= nc.S[i][1] {
			k = k - nc.S[i][1]
			ans--
			i++
		} else {
			break
		}
	}
	return ans
}
