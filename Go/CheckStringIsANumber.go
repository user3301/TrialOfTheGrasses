package gosoln

func IsNumberString(s string) bool {
	stateMachine := []map[rune]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4},
		{'d': 2, '.': 4},
		{'d': 2, '.': 3, 'e': 5, ' ': 8},
		{'d': 3, 'e': 5, ' ': 8},
		{'d': 3},
		{'s': 6, 'd': 7},
		{'d': 7},
		{'d': 7, ' ': 8},
		{' ': 8},
	}
	p := 0
	var t rune
	for _, chr := range s {
		if chr >= '0' && chr <= '9' {
			t = 'd'
		} else if chr == '+' || chr == '-' {
			t = 's'
		} else if chr == 'e' || chr == 'E' {
			t = 'e'
		} else if chr == '.' || chr == ' ' {
			t = chr
		} else {
			t = '?'
		}
		if _, ok := stateMachine[p][t]; !ok {
			return false
		}
		p = stateMachine[p][t]

	}
	return p == 2 || p == 3 || p == 7 || p == 8
}
