package gosoln

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
