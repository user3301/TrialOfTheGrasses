package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestTheKWeakestRows1(t *testing.T) {
	input := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	k := 3
	expected := []int{2, 0, 3}
	actual := KWeakestRows(input, k)
	assert.DeepEqual(t, expected, actual)
}

func TestTheKWeakestRows2(t *testing.T) {
	input := [][]int{
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 1},
	}
	k := 4
	expected := []int{0, 1, 2, 3}
	actual := KWeakestRows(input, k)
	assert.DeepEqual(t, expected, actual)
}

func TestTheKWeakestRows3(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	k := 1
	expected := []int{0}
	actual := KWeakestRows(input, k)
	assert.DeepEqual(t, expected, actual)
}

func TestTheKWeakestRows4(t *testing.T) {
	input := [][]int{
		{1, 0},
		{0, 0},
		{1, 0},
	}
	k := 2
	expected := []int{1, 0}
	actual := KWeakestRows(input, k)
	assert.DeepEqual(t, expected, actual)
}
