package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestClosestDivisors1(t *testing.T) {
	input := 123
	expected := []int{5, 25}
	actual := ClosestDivisors(input)
	assert.DeepEqual(t, expected, actual)
}
