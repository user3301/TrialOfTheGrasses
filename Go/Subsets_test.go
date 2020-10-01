package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	want := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	got := Subsets(nums)
	assert.DeepEqual(t, want, got)
}
