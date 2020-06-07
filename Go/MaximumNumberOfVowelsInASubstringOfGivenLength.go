package gosoln

type void struct{}

func MaxVowels(s string, k int) int {
	vowels := map[uint8]struct{}{'a': void{}, 'e': void{}, 'i': void{}, 'o': void{}, 'u': void{}}
	ans, cur := 0, 0
	left, right := 0, 0
	for right < len(s) {
		_, isv := vowels[s[right]]
		if isv {
			cur++
		}
		if right-left >= k {
			_, isvl := vowels[s[left]]
			if isvl {
				cur--
				left++
			} else {
				left++
			}
		}
		right++
		ans = max(ans, cur)
	}
	return ans
}
