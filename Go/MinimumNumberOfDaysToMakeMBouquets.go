package gosoln

func MinDays(bloomDays []int, m, k int) int {
	ans := -1
	if m*k > len(bloomDays) {
		return ans
	}
	left, right := 0, int(^uint(0)>>1)
	for left <= right {
		mid := left + (right-left)/2
		if canHarvest(bloomDays, m, k, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func canHarvest(bd []int, m, k, mid int) bool {
	bouq, flower := 0, 0
	for i := 0; i < len(bd); i++ {
		if bd[i] <= mid {
			flower++
			if flower == k {
				bouq++
				flower = 0
			}
		} else {
			flower = 0
		}
	}
	return bouq >= m
}
