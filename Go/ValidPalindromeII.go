package Go

func ValidPalindromeII(s string) bool {
	re := []rune(s)

	l, r := 0, len(s)-1
	for l < r {
		if re[l] == re[r] {
			l++
			r--
		} else {
			return isPalindrome(re[l+1:r+1]) || isPalindrome(re[l:r])
		}
	}
	return true
}

func isPalindrome(re []rune) bool {
	l, r := 0, len(re)-1

	for l < r {
		if re[l] == re[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
