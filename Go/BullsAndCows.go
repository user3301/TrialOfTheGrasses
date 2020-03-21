package Go

import (
	"strconv"
)

func GetHint(secret string, guess string) string {
	bulls, cows := 0, 0
	numbers := [10]int{}

	r1 := []rune(secret)
	r2 := []rune(guess)

	for i := 0; i < len(r1); i++ {
		if r1[i] == r2[i] {
			bulls++
		} else {
			r1int, _ := strconv.Atoi(string(r1[i]))
			r2int, _ := strconv.Atoi(string(r2[i]))
			if numbers[r1int] < 0 {
				cows++
			}
			if numbers[r2int] > 0 {
				cows++
			}
			numbers[r1int]++
			numbers[r2int]--
		}
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
