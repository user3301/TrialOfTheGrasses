package gosoln

func SubtractProductAndSum(n int) int {
	sum, prod := 0, 1

	for n != 0 {
		digit := n % 10
		sum += digit
		prod *= digit
		n /= 10
	}

	return prod - sum
}
