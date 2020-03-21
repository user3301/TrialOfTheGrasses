package Go

func ReverseInteger(x int) int {
	if x-int(int32(x)) > 0 {
		return 0
	}

	p := x
	if x < 0 {
		p = -p
	}
	arr := []int{}

	d := 10

	for p > 0 {
		arr = append(arr, p%10)
		p /= 10
		d *= 10
	}

	ans := 0
	d /= 100

	for i := 0; i < len(arr); i++ {
		ans += arr[i] * d
		if ans != int(int32(ans)) {
			return 0
		}
		d /= 10
	}
	if x < 0 {
		if -ans != int(int32(-ans)) {
			return 0
		}
		return -ans
	}

	return ans
}
