package Go

func Shift2DGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < n*m; i++ {
		j := (i + k) % (m * n)
		ans[j/m][j%m] = grid[i/m][i%m]
	}
	return ans
}
