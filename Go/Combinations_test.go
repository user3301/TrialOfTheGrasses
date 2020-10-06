package gosoln

import (
	"gotest.tools/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	n, k := 4, 2
	want := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	got := Combine(n, k)
	assert.DeepEqual(t, want, got)
}
