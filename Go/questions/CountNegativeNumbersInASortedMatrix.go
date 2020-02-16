package questions

// CountNegativesBF brute force solution O(n^2)
// Loop through all elements and check them
func CountNegativesBF(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}

// CountNegativesBFOptimized optimized brute force solution
// O(n^2) worst case, counting backwards, early termination when element >0
func CountNegativesBFOptimized(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] >= 0 {
				break
			}
			count++
		}
	}
	return count
}

// CountNegativesLinear complexity O(m +n)
//Start from bottom-left corner of the matrix, count in the negative numers in each row.
func CountNegativesLinear(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	r, c := m-1, 0
	count := 0

	for r >= 0 && c < n {
		if grid[r][c] < 0 {
			r--
			count += n - c
		} else {
			c++
		}
	}
	return count
}
