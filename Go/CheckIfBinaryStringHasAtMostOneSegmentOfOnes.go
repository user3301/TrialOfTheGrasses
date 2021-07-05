package gosoln

func checkOnesSegment(s string) bool {
	cnt := 0
	var chr = '1'
	for _, c := range s {
		if c^chr == 1 {
			cnt++
			chr = c
			if cnt > 1 {
				return false
			}
		}
	}
	return true
}
