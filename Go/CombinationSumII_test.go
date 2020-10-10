package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSumII(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	want := [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}
	got := CombinationSum2(candidates, target)
	assert.ElementsMatch(t, want, got)
}
