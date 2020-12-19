package gosoln

func NumberOfMatches(n int) int {
	var matchesPlayed int

	for n >= 2 {
		if n%2 == 0 {
			// even
			matchesPlayed += n / 2
			n /= 2
		} else {
			// odd
			matchesPlayed += n / 2
			n = 1 + (n / 2)
		}
	}
	return matchesPlayed
}
