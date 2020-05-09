package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	actual := MaxArea(input)
	assert.Equal(t, expected, actual)
}
