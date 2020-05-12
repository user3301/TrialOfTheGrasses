package gosoln

func CountTriplets(arr []int) int {
	ans := 0
	prefix := make([]int, len(arr)+1)

	for i := 1; i < len(prefix); i++ {
		prefix[i] = arr[i-1] ^ prefix[i-1]
	}
	for i := 0; i < len(prefix); i++ {
		for j := i + 1; j < len(prefix); j++ {
			for k := j + 1; k < len(prefix); k++ {
				if prefix[j]^prefix[i] == (prefix[k] ^ prefix[j]) {
					ans++
				}
			}
		}
	}
	return ans
}
