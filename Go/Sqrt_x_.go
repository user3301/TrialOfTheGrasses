package Go

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}

	l, r := 1, x

	for l+1 < r {
		mid := (l + r) >> 1
		if mid*mid > x {
			r = mid
		} else if mid*mid < x {
			l = mid
		} else {
			return mid
		}
	}
	return l
}
