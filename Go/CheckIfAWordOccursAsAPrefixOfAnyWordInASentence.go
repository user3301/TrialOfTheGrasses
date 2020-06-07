package gosoln

import "strings"

func IsPrefixOfWord(sentence, searchWord string) int {
	ans := -1
	wordsArr := strings.Split(sentence, " ")
	for i, w := range wordsArr {
		if isPrefix(w, searchWord) {
			ans = i + 1
			return ans
		}
	}
	return ans
}

func isPrefix(target, prefix string) bool {
	if len(target) < len(prefix) {
		return false
	}
	for i := range prefix {
		if target[i] != prefix[i] {
			return false
		}
	}
	return true
}
