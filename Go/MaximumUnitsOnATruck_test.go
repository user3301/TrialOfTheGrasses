package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaximumUnits(t *testing.T) {
	boxTypes := [][]int{
		{1, 3}, {2, 2}, {3, 1},
	}
	got := MaximumUnits(boxTypes, 4)
	want := 8
	assert.Equal(t, want, got)
}

func TestMaximumUnits2(t *testing.T) {
	boxTypes := [][]int{
		{5, 10}, {2, 5}, {4, 7}, {3, 9},
	}
	got := MaximumUnits(boxTypes, 10)
	want := 91
	assert.Equal(t, want, got)
}
