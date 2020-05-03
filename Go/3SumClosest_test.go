package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func Test3SumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	expected := 2
	actual := ThreeSumClosest(nums, target)
	assert.Equal(t, expected, actual)
}
