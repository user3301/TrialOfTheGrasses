package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinTime1(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	hasApple := []bool{false, false, true, false, true, true, false}
	expected := 8
	actual := MinTime(n, edges, hasApple)
	assert.Equal(t, expected, actual)
}

func TestMinTime2(t *testing.T) {
	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	hasApple := []bool{false, false, true, false, false, true, false}
	expected := 6
	actual := MinTime(n, edges, hasApple)
	assert.Equal(t, expected, actual)
}
