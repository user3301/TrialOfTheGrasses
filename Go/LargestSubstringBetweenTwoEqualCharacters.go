package gosoln

import "github.com/user3301/TrialOfTheGrasses/Go/pkg/utils"

func MaxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune]int)
	ans := -1
	for i, v := range s {
		if _, ok := m[v]; ok {
			distance := i - m[v] - 1
			ans = utils.MaxInt(ans, distance)
		} else {
			m[v] = i
		}
	}
	return ans
}
