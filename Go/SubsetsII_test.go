package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets2(t *testing.T) {
	input := []int{1, 2, 2}
	want := make([][]int, 6)
	want[0] = nil
	want[1] = []int{1}
	want[2] = []int{1, 2}
	want[3] = []int{1, 2, 2}
	want[4] = []int{2}
	want[5] = []int{2, 2}
	//want := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	got := SubsetsWithDup(input)
	assert.ElementsMatch(t, want, got)
}
