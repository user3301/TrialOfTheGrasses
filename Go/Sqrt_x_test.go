package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestSqrt_x_1(t *testing.T) {
	input := 4
	expected := 2
	actual := MySqrt(input)
	assert.Equal(t, expected, actual)
}

func TestSqrt_x_2(t *testing.T) {
	input := 8
	expected := 2
	actual := MySqrt(input)
	assert.Equal(t, expected, actual)
}
