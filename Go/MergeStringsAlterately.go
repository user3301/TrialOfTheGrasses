package gosoln

//func MergeAlternately(word1 string, word2 string) string {
//	l1, l2 := len(word1), len(word2)
//	final := make([]byte, l1+l2)
//	i := 0
//	j := 0
//	for i < l1 && i < l2 {
//		final[j] = word1[i]
//		j++
//		final[j] = word2[i]
//		j++
//		i++
//	}
//
//	if l1 > l2 {
//		for i < l1 {
//			final[j] = word1[i]
//			j++
//			i++
//		}
//	} else {
//		for i < l2 {
//			final[j] = word2[i]
//			j++
//			i++
//		}
//	}
//	return string(final)
//}

func mergeAlternately(word1, word2 string) string {
	l1, l2, i, j := len(word1), len(word2), 0, 0
	ans := make([]byte, 0)
	for i < l1 || j < l2 {
		if i < l1 {
			ans = append(ans, word1[i])
			i++
		}
		if j < l2 {
			ans = append(ans, word2[j])
			j++
		}
	}
	return string(ans)
}
