package Go

func FindPairs(nums []int, k int) int {
	ans := 0
	dist := make(map[int]int)
	for _, v := range nums {
		dist[v]++
	}
	for i, v := range dist {
		_, ok := dist[i+k]
		if k > 0 && ok || k == 0 && v >= 2 {
			ans++
		}
	}
	return ans
}
