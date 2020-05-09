package gosoln

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}

	l, r := 1, x

	for l+1 < r {
		mid := (l + r) >> 1
		switch {
		case mid*mid > x:
			r = mid
		case mid*mid < x:
			l = mid
		default:
			return mid
		}
	}
	return l
}
