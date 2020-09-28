package gosoln

func ReorderSpaces(text string) string {
	var words []string
	spaceCnt := 0
	var tempWord []rune
	for _, rv := range text {
		if rv == ' ' {
			spaceCnt++
			if len(tempWord) != 0 {
				words = append(words, string(tempWord))
				tempWord = []rune{}
			}
		} else {
			tempWord = append(tempWord, rv)
		}
	}
	if len(tempWord) != 0 {
		words = append(words, string(tempWord))
	}
	spaceBetween, remain := 0, 0
	if len(words) == 1 {
		remain = spaceCnt
	} else {
		spaceBetween = spaceCnt / (len(words) - 1)
		remain = spaceCnt % (len(words) - 1)
	}

	ans := ""
	for j, w := range words {
		ans += w
		for i := 0; i < spaceBetween && j < len(words)-1; i++ {
			ans += " "
		}

	}
	for i := 0; i < remain; i++ {
		ans += " "
	}
	return ans
}
