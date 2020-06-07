package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxArea1(t *testing.T) {
	h, w := 5, 4
	horizontalCuts := []int{1, 2, 4}
	verticalCuts := []int{1, 3}
	expected := 4
	actual := MaxArea1465(h, w, horizontalCuts, verticalCuts)
	assert.Equal(t, expected, actual)
}
