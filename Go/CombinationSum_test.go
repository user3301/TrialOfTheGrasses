package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	want := [][]int{{2, 2, 3}, {7}}
	got := CombinationSum(candidates, target)
	assert.ElementsMatch(t, want, got)
}
