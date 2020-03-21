package Go

func ReverseVowels(s string) string {
	srune := []rune(s)
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	left, right := 0, len(s)-1
	for left < right {
		for {
			_, lok := vowels[srune[left]]
			if left < right && !lok {
				left++
			} else {
				break
			}
		}
		for {
			_, rok := vowels[srune[right]]
			if left < right && !rok {
				right--
			} else {
				break
			}
		}
		srune[left], srune[right] = srune[right], srune[left]
		left++
		right--
	}
	return string(srune)
}
