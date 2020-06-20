package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestShuffle(t *testing.T) {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	expected := []int{2, 3, 5, 4, 1, 7}
	actual := Shuffle(nums, n)
	assert.DeepEqual(t, expected, actual)
}
