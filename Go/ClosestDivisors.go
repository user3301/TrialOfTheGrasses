package Go

import "math"

func ClosestDivisors(num int) []int {
	ans := make([]int, 0, 2)
	min := int(^uint(0) >> 1)

	first := num + 1
	mid := int(math.Sqrt(float64(first)))

	for i := mid + 1; i > 0; i-- {
		if first%i == 0 {
			diff := abs(i - (first / i))
			if diff < min {
				min = diff
				ans = make([]int, 0, 2)
				ans = append(ans, i)
				ans = append(ans, first/i)
				break
			}
		}
	}

	second := num + 2
	mid = int(math.Sqrt(float64(second)))
	for i := mid + 1; i > 0; i-- {
		if second%i == 0 {
			diff := abs(i - (second / i))
			if diff < min {

				ans = make([]int, 0, 2)
				ans = append(ans, i)
				ans = append(ans, second/i)
				break
			}
		}
	}
	return ans
}
