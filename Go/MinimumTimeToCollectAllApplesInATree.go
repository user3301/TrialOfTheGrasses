package gosoln

func MinTime(n int, edges [][]int, hasApple []bool) int {
	ans := 0
	for i := len(edges) - 1; i >= 0; i-- {
		if hasApple[edges[i][1]] {
			hasApple[edges[i][0]] = true
		}
	}
	for i := 0; i < len(edges); i++ {
		if hasApple[edges[i][1]] {
			ans++
		}
	}
	return ans * 2
}
