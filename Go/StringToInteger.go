package gosoln

import (
	"math"
	"unicode"
)

func MyAtoi(s string) int {
	stateMachine := map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	var switchState func(s rune) int
	switchState = func(s rune) int {
		switch s {
		case ' ':
			return 0
		case '+', '-':
			return 1
		default:
			if unicode.IsDigit(s) {
				return 2
			}
		}
		return 3
	}
	state := "start"
	sign := 1
	ans := 0

	for _, v := range s {
		state = stateMachine[state][switchState(v)]
		switch state {
		case "start":
			continue
		case "in_number":
			ans = ans*10 + ToNumber(v)
			if sign > 0 && ans*sign > math.MaxInt32 {
				return math.MaxInt32
			}
			if sign < 0 && ans*sign < math.MinInt32 {
				return math.MinInt32
			}
		case "signed":
			if v == '-' {
				sign = -1
			}
		case "end":
			break
		}
	}
	return sign * ans
}

func ToNumber(r rune) int {
	return int(r - '0')
}
