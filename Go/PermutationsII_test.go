package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestPermuteUnique(t *testing.T) {
	input := []int{1, 1, 2}
	want := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
	got := PermuteUnique(input)
	assert.DeepEqual(t, want, got)
}
