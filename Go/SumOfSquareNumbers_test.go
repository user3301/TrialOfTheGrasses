package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestSumOfSquareNumbers1(t *testing.T) {
	input := 0
	actual := JudgeSquareSum(input)
	assert.Equal(t, true, actual)
}

func TestSumOfSquareNumbers2(t *testing.T) {
	input := 5
	actual := JudgeSquareSum(input)
	assert.Equal(t, true, actual)
}

func TestSumOfSquareNumbers3(t *testing.T) {
	input := 3
	actual := JudgeSquareSum(input)
	assert.Equal(t, false, actual)
}

func TestSumOfSquareNumbers4(t *testing.T) {
	input := 4
	actual := JudgeSquareSum(input)
	assert.Equal(t, true, actual)
}
