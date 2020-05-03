package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestXorQueries(t *testing.T) {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
		{3, 3},
	}
	expected := []int{2, 7, 14, 8}
	actual := XorQueries(arr, queries)
	assert.DeepEqual(t, expected, actual)
}
