package gosoln

func XorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	preTable := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		preTable[i+1] = preTable[i] ^ arr[i]
	}

	for i := 0; i < len(queries); i++ {
		l := queries[i][0]
		r := queries[i][1]
		ans[i] = preTable[r+1] ^ preTable[l]
	}
	return ans
}
