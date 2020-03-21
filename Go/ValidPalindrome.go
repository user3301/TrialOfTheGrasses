package Go

import "unicode"

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	r := []rune(s)
	for left < right {
		for left < right && !isLetterOrDigit(unicode.ToLower(r[left])) {
			left++
		}
		for right > left && !isLetterOrDigit(unicode.ToLower(r[right])) {
			right--
		}
		if unicode.ToLower(r[left]) != unicode.ToLower(r[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrDigit(r rune) bool {
	return ('a' <= r && r <= 'z') || ('0' <= r && r <= '9')
}
