package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"testing"
)

func TestTheKWeakestRows(t *testing.T) {
	input := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	k := 3
	expected := []int{2, 0, 3}
	actual := questions.KWeakestRows(input, k)
	assert.DeepEqual(t, expected, actual)
}
