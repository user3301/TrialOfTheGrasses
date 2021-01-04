package gosoln

import "sort"

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	var ans int
	bu := &ByUnits{boxTypes}
	sort.Stable(*bu)
	i, j := len(boxTypes), 0

	for j < i && truckSize > 0 {
		if truckSize >= bu.boxes[j][0] {
			truckSize -= bu.boxes[j][0]
			ans += bu.boxes[j][0] * bu.boxes[j][1]
			j++
		} else {
			ans += bu.boxes[j][1] * truckSize
			truckSize = 0
		}
	}
	return ans
}

type ByUnits struct {
	boxes [][]int
}

func (b ByUnits) Len() int {
	return len(b.boxes)
}

func (b ByUnits) Swap(i, j int) {
	b.boxes[i], b.boxes[j] = b.boxes[j], b.boxes[i]
}

func (b ByUnits) Less(i, j int) bool {
	return b.boxes[j][1] < b.boxes[i][1]
}
