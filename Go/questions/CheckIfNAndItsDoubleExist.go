package questions

func CheckIfExist(arr []int) bool {
	seen := map[int]bool{}
	for _, k := range arr {
		if seen[k*2] {
			return true
		}
		if k%2 == 0 && seen[k/2] {
			return true
		}
		seen[k] = true
	}
	return false
}
