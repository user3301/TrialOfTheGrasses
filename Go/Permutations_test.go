package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestPermute1(t *testing.T) {
	arr := []int{1, 2, 3}
	got := Permute(arr)
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	assert.DeepEqual(t, want, got)
}

func TestPermute2(t *testing.T) {
	arr := []int{5, 4, 6, 2}
	got := Permute(arr)
	want := [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2},
		{4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2},
		{6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4},
		{2, 6, 4, 5}}
	assert.DeepEqual(t, want, got)
}
