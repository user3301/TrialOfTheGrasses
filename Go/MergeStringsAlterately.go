package gosoln

func MergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	final := make([]byte, l1+l2)
	i := 0
	j := 0
	for i < l1 && i < l2 {
		final[j] = word1[i]
		j++
		final[j] = word2[i]
		j++
		i++
	}

	if l1 > l2 {
		for i < l1 {
			final[j] = word1[i]
			j++
			i++
		}
	} else {
		for i < l2 {
			final[j] = word2[i]
			j++
			i++
		}
	}
	return string(final)
}
