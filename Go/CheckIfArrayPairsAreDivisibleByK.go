package gosoln

func CanArrange(arr []int, k int) bool {
	reminder := make(map[int]int)
	for _, v := range arr {
		reminder[(v%k+k)%k]++
	}
	for key := range reminder {
		if key == 0 {
			continue
		}
		if _, exist := reminder[k-key]; !exist || reminder[k-key] != reminder[key] {
			return false
		}
	}
	return reminder[0]%2 == 0
}
