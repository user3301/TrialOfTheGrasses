package gosoln

func MinPartitions(n string) int {
	strLen := len(n)
	var finalRs []uint8
	for i := 0; i < strLen; i++ {
		finalRs = append(finalRs, '0')
	}
	str := string(finalRs)

	rs := []rune(n)
	var ret int
	for string(rs) != str {
		for i := 0; i < strLen; i++ {
			if rs[i] != '0' {
				rs[i] -= 1
			}
		}
		ret++
	}
	return ret
}
