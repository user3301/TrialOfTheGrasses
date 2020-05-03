package gosoln

func BackspaceCompare(s string, t string) bool {
	srunes := []rune(s)
	trunes := []rune(t)
	sback, tback, sr, tr := len(srunes)-1, len(trunes)-1, 0, 0

	for sback >= 0 || tback >= 0 {
		for sback >= 0 && (srunes[sback] == '#' || sr > 0) {
			if srunes[sback] == '#' {
				sr++
			} else {
				sr--
			}
			sback--
		}
		for tback >= 0 && (trunes[tback] == '#' || tr > 0) {
			if tback >= 0 && (trunes[tback] == '#' || tr > 0) {
				if trunes[tback] == '#' {
					tr++
				} else {
					tr--
				}
			}
			tback--
		}
		if sback < 0 || tback < 0 {
			return sback == tback
		}
		if srunes[sback] != trunes[tback] {
			return false
		}
		sback--
		tback--
	}
	return true

}
