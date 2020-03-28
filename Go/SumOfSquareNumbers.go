package Go

import "math"

func JudgeSquareSum(c int) bool {
	sqrt := math.Sqrt(float64(c))
	return canFind(sqrt, c)
}

func canFind(f float64, c int) bool {

	l, r := 0, int(f)

	for l <= r {
		switch {
		case l*l+r*r == c:
			return true
		case l*l+r*r > c:
			r--
		default:
			l++
		}
	}
	return false
}
