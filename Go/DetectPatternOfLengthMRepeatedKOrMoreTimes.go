package gosoln

func ContainasPattern(arr []int, m int, k int) bool {
	cnt := 0
	for i := 0; i+m < len(arr); i++ {
		if arr[i] != arr[i+m] {
			cnt = 0
		} else {
			cnt++
		}
		if cnt == (k-1)*m {
			return true
		}
	}
	return false
}
