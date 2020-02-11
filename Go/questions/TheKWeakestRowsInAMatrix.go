package questions

func KWeakestRows(mat [][]int, k int) []int {
	bucket := make([][]int, len(mat[0])+1)
	for i := 0; i < len(mat); i++ {
		bucket[i] = make([]int, 0)
	}
	for i := 0; i < len(mat); i++ {
		count :=0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] ==1 {
				count++
			} else {
				break
			}
		}
		bucket[count] = append(bucket[count], i)
	}

	ans := []int{}

	for i:= 0; i < len(bucket); i++ {
		for len(bucket[i]) > 0 && k > 0 {
			ans = append(ans, bucket[i][0])
			bucket[i] = bucket[i][1:]
			k--
		}
	}
	return ans
}
