package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxProduct(t *testing.T) {
	input := []int{3, 4, 5, 2}
	expected := 12
	actual := MaxProduct(input)
	assert.Equal(t, expected, actual)
}
