package questions

func NumberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}
