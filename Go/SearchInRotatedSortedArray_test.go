package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4
	actual := SearchInRotatedSortedArray(nums, target)
	assert.Equal(t, expected, actual)
}
